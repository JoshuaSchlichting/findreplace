package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
)

var testFiles []string

const (
	findPhrase  = "replacethis"
	replaceWith = "success"
	testDataDir = "/tmp/testdata/"
)

// TestMain is a low-level primitive and should not be necessary for casual testing needs,
// where ordinary test functions suffice.
// https://pkg.go.dev/testing#hdr-Main
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	if _, err := os.Stat(testDataDir); os.IsNotExist(err) {
		err := os.Mkdir(testDataDir, 0777)
		fmt.Println(fmt.Sprintf("An error occurred while attempting to create the test data directory: '%s'"+
			testDataDir, err))
	}
}

func tearDown() {
	os.Remove(testDataDir)
}

func handleFileError(filename string, err error) {
	if err != nil {
		errMsg := fmt.Sprintf("An error occurred while reading %s\n%s", filename, err)
		fmt.Println(errMsg)
	}
}

func TestFindReplaceASingleFile(t *testing.T) {
	testFilename := generateTestFile(t, string(findPhrase))
	findAndReplace(testFilename, findPhrase, replaceWith)
	fileText, err := ioutil.ReadFile(testFilename)
	handleFileError(testFilename, err)
	if strings.Contains(string(fileText), string(findPhrase)) {
		t.Errorf("The find phrase '%s' was not replaced in the test file '%s'", findPhrase, testFilename)
	}
}

func TestFindReplaceDirectory(t *testing.T) {
	loopCount := 2
	for loopCount > 0 {
		testFiles = append(testFiles, generateTestFile(t, findPhrase))
		loopCount--
	}

	findAndReplace(testDataDir, findPhrase, replaceWith)
	for _, filename := range testFiles {
		fileText, err := ioutil.ReadFile(filename)
		handleFileError(filename, err)
		if strings.Contains(string(fileText), string(findPhrase)) {
			t.Errorf("The find phrase '%s' was not replaced in the test file '%s'", findPhrase, filename)
		}
	}
}

func generateTestFile(t *testing.T, seedString string) string {

	fmt.Println("Generating test file...")
	testFilename := testDataDir + "unittest_data_" + strconv.Itoa(rand.Int())
	err := ioutil.WriteFile(testFilename, []byte(seedString), 0644)
	if err != nil {
		fmt.Println(err)
	}

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	return testFilename
}
