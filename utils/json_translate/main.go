package main

import (
	"bytes"
	"encoding/json"
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

type anyJSONObject interface{}

type orderedMap struct {
	Value interface{}
	Keys  []string
}

func (um orderedMap) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &um.Value); err != nil {
		return err
	}
	fmt.Println("...")
	for key, _ := range um.Value.(map[string]interface{}) {
		fmt.Println("---", key)
		um.Keys = append(um.Keys, key)
	}
	return nil
}

func (um orderedMap) MarshalJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	buf.WriteRune('{')
	l := len(um.Keys)
	if um.Value == nil {
		buf.WriteRune('}')
		return buf.Bytes(), nil
	}
	valueMap := um.Value.(map[string]interface{})
	for i, key := range um.Keys {
		km, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(km)
		buf.WriteRune(':')
		vm, err := json.Marshal(valueMap[key])
		if err != nil {
			return nil, err
		}
		buf.Write(vm)
		if i != l-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune('}')
	return buf.Bytes(), nil
}

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

	helpers.SetTranslateConfig(helpers.TranslateConfig{LanguageFrom: lFront, LanguageTo: lTo})

	filePaths, err := helpers.GetFilesFromDir(path.Join(currentPath, FILES_PATH), "*.json")
	if err != nil {
		fmt.Printf("Error %v \n", err)
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
