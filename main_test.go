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

func TestFindReplaceASingleFile(t *testing.T) {
	testFilename := generateTestFile(string(findPhrase))
	findAndReplace(testFilename, findPhrase, replaceWith)
	fileText, err := ioutil.ReadFile(testFilename)
	if err != nil {
		fmt.Printf("An error occurred while reading %s\n%s", testFilename, err)
	}

	if strings.Contains(string(fileText), string(findPhrase)) {
		t.Errorf("The find phrase '%s' was not replaced in the test file '%s'", findPhrase, testFilename)
	}
}

func TestFindReplaceDirectory(t *testing.T) {
	loopCount := 5
	for loopCount > 0 {
		append(testFiles, generateTestFile(findPhrase))
		loopCount--
	}
	for _, filename := range testFiles {
		findAndReplace(filename, findPhrase, replaceWith)
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
