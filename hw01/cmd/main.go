package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const extSeparator = "."

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	fileName, fileExt := getFileNameAndExt(filePth, string(os.PathSeparator))

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}

func getFileNameAndExt(filePath, sep string) (string, string) {
	splittedPath := strings.Split(filePath, sep)
	base := splittedPath[len(splittedPath) - 1]
	splittedBase := strings.Split(base, extSeparator)
	lastIdx := len(splittedBase) - 1

	ext := splittedBase[lastIdx]
	filename := strings.Join(splittedBase[:lastIdx], extSeparator)

	if filename == "" {
		return base, ""
	}

	return filename, ext
}