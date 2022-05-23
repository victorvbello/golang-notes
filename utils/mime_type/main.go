package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func GetFilesFromDir(root string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		matches = append(matches, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func GetFileContentType(file *os.File) (string, error) {
	buffer := make([]byte, 512)

	if _, err := file.Read(buffer); err != nil && err != io.EOF {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func main() {
	filePaths, err := GetFilesFromDir("./files")
	if err != nil {
		fmt.Printf("Error %v \n", err)
	}
	for fileIndex, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("os.Open %d Error %v\n", fileIndex, err)
			continue
		}
		mime, err := GetFileContentType(file)
		if err != nil {
			fmt.Printf("GetFileContentType %d Error %v\n", fileIndex, err)
			continue
		}
		fmt.Printf("[%d]\t%s\t%s\n", fileIndex, filePath, mime)
	}
}
