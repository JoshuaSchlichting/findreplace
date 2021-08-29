package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var testFiles []string

const findPhrase = "replacethis"
const replaceWith = "success"

func handleFileError(filename string, err error) {
	if err != nil {
		errMsg := fmt.Sprintf("An error occurred while reading %s\n%s", filename, err)
		fmt.Println(errMsg)
	}
}

func TestFindReplaceASingleFile(t *testing.T) {
	testFilename := generateTestFile(string(findPhrase))
	findAndReplace(testFilename, findPhrase, replaceWith)
	fileText, err := ioutil.ReadFile(testFilename)
	handleFileError(testFilename, err)
	if strings.Contains(string(fileText), string(findPhrase)) {
		t.Errorf("The find phrase '%s' was not replaced in the test file '%s'", findPhrase, testFilename)
	}
}

func TestFindReplaceDirectory(t *testing.T) {
	loopCount := 5
	for loopCount > 0 {
		testFiles = append(testFiles, generateTestFile(findPhrase))
		loopCount--
	}
	for _, filename := range testFiles {
		findAndReplace(filename, findPhrase, replaceWith)
		fileText, err := ioutil.ReadFile(filename)
		handleFileError(filename, err)
		if strings.Contains(string(fileText), string(findPhrase)) {
			t.Errorf("The find phrase '%s' was not replaced in the test file '%s'", findPhrase, filename)
		}
	}
}

func generateTestFile(seedString string) string {
	fmt.Println("Generating test file...")
	testFilename := "/tmp/unittest_data_" + strconv.Itoa(rand.Int())
	err := ioutil.WriteFile(testFilename, []byte(seedString), 0644)
	if err != nil {
		fmt.Println(err)
	}
	return testFilename
}
