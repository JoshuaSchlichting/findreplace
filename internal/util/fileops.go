package util

import (
	"fmt"
	"os"
)

func GetFileStatMode(filePath string) os.FileMode {
	fileStat, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fileStat.Mode()
}
