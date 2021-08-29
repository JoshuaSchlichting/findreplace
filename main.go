package main

import (
	"findreplace/internal/util"
	"flag"
	"fmt"
)

func main() {
	pathToSearch := flag.String("path", "", "Path to target for find-and-replace. NOTE: This can be a single filename or a directory. (Required)")
	findPhrase := flag.String("find", "", "String to find.")
	replacePhrase := flag.String("replace", "", "String to replace.")
	flag.Parse()

	findAndReplace(*pathToSearch, *findPhrase, *replacePhrase)
}

func findAndReplace(pathToSearch, findPhrase, replacePhrase string) {
	switch mode := util.GetFileStatMode(pathToSearch); {
	case mode.IsDir():
		// do directory stuff
		fmt.Println("directory")
		util.ReplacePhrasesInDirectory(pathToSearch, findPhrase, replacePhrase)
	case mode.IsRegular():
		fmt.Println("file")
		util.FindAndReplace(pathToSearch, findPhrase, replacePhrase)
	}

}
