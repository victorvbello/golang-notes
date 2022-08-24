package main

import (
	"fmt"
	"gonotes/utils/json_translate/helpers"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const (
	FILES_PATH        = "./files"
	FILES_OUTPUT_PATH = "./output"
)

func main() {
	args := os.Args[1:]
	var currentPath, lFront, lTo string

	if len(os.Args) < 3 {
		log.Fatal("path, language from and language to are required")
	}

	currentPath = args[0]
	lFront = args[1]
	lTo = args[2]

	if currentPath == "" || lFront == "" || lTo == "" {
		log.Fatal("path, language from and language to are required")
	}

	currentInputPath := path.Join(currentPath, FILES_PATH)

	helpers.SetTranslateConfig(helpers.TranslateConfig{LanguageFrom: lFront, LanguageTo: lTo})

	filePaths, err := helpers.GetFilesFromDir(currentInputPath, "*.json")
	if err != nil {
		fmt.Printf("Error %v \n", err)
	}
	if len(filePaths) == 0 {
		fmt.Printf("Error without input .json files in %s\n", currentInputPath)
	}
	for fileIndex, filePath := range filePaths {
		start := time.Now()
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("os.Open %d Error %v\n", fileIndex, err)
			continue
		}
		fileName := filepath.Base(file.Name())
		fileNameClean := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		fmt.Printf("Process [%s]\n", fileNameClean)
		jsonContent, err := helpers.GetFileContent(file)
		if err != nil {
			fmt.Printf("GetFileContent %d Error %v\n", fileIndex, err)
			continue
		}

		finalJson, err := helpers.GetMapValue(jsonContent)
		if err != nil {
			fmt.Printf("GetMapValue %d Error %v\n", fileIndex, err)
			continue
		}

		f, err := os.Create(path.Join(currentPath, FILES_OUTPUT_PATH, fileNameClean) + "_output.json")
		if err != nil {
			fmt.Printf("os.Create %d Error %v\n", fileIndex, err)
			continue
		}
		defer f.Close()
		_, err = f.Write(finalJson)
		if err != nil {
			fmt.Printf("f.Write %d Error %v\n", fileIndex, err)
			continue
		}
		elapsed := time.Since(start)
		fmt.Printf("End [%s] Translates: %d In: %s \n", fileNameClean, helpers.TOTAL_TRANSLATE, elapsed)
		helpers.TOTAL_TRANSLATE = 0
	}
}
