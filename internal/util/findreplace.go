package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func ReplacePhrasesInDirectory(dirPath, find, replace string) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// rm current dir from files before loop
	for _, filePath := range files[1:] {
		fmt.Printf("Checking %q\n", filePath)

		switch mode := GetFileStatMode(filePath); {
		case mode.IsDir():
			fmt.Printf("Checking %q for text files...\n", filePath)
			ReplacePhrasesInDirectory(filePath, find, replace)
		case mode.IsRegular():
			mtype, err := mimetype.DetectFile(filePath)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%q is of mime-type %q", filePath, mtype.String())
			if strings.Contains(mtype.String(), "text/plain") {
				FindAndReplace(filePath, find, replace)
			} else {
				fmt.Printf("%q is not a text file. Skipping...\n", filePath)
			}
		}
	}
}

func FindAndReplace(filePath, find, replace string) {
	fmt.Printf("Searching for %s in %s\n", find, filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	newFileContent := strings.Replace(string(data[:]), find, replace, -1)
	if err := os.Truncate(filePath, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filePath, []byte(newFileContent), 0644); err != nil {
		log.Printf("Unable to write data to file!")
		os.Exit(1)
	}
	fmt.Println("Contents of file:", string(data))
}
