package main

import (
	"encoding/base64"
	"fmt"
	"io"
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
		defer file.Close()
		b, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("io.ReadAll %d Error %v\n", fileIndex, err)
			continue
		}

		b64Enc := base64.StdEncoding.EncodeToString(b)

		fmt.Printf("[%d]\t%s\t%s\n", fileIndex, filePath, b64Enc)
	}
}
