package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gonotes/utils/json_translate/orderedjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

const (
	TRANSLATE_API_URL = "http://localhost:5000/translate"
)

type requestResult map[string]interface{}
type requestHeaders map[string]string
type TranslateConfig struct {
	LanguageFrom string
	LanguageTo   string
}

var TOTAL_TRANSLATE = 0
var translate_config *TranslateConfig

func SetTranslateConfig(tc TranslateConfig) {
	translate_config = &tc
}

func GetTranslateConfig() *TranslateConfig {
	return translate_config
}

func HttpRequest(method string, url string, headers requestHeaders, body []byte) (requestResult, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	var resultBody requestResult
	var finalBody io.Reader
	if len(body) > 0 {
		finalBody = bytes.NewBuffer(body)
	}
	req, err := http.NewRequest(method, url, finalBody)
	if err != nil {
		return resultBody, fmt.Errorf("http.NewRequest %w", err)
	}
	if len(headers) > 0 {
		for key, data := range headers {
			req.Header.Set(key, data)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return resultBody, fmt.Errorf("client.Do %w", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resultBody, fmt.Errorf("ioutil.ReadAll %w", err)
	}
	err = json.Unmarshal(body, &resultBody)
	if err != nil {
		return resultBody, fmt.Errorf("json.Unmarshal %w", err)
	}
	return resultBody, nil
}

func GetFilesFromDir(root string, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func GetFileContent(file *os.File) ([]byte, error) {
	fileInfo, err := file.Stat()

	if err != nil {
		return []byte{}, err
	}

	buffer := make([]byte, fileInfo.Size())

	if _, err := file.Read(buffer); err != nil && err != io.EOF {
		return []byte{}, err
	}

	return buffer, nil
}

func TranslateString(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	headers := requestHeaders(map[string]string{
		"content-type": "application/json",
	})
	bodyTranslate := map[string]interface{}{
		"q":      s,
		"source": GetTranslateConfig().LanguageFrom,
		"target": GetTranslateConfig().LanguageTo,
		"format": "text",
	}
	bodyTranslateBytes, err := json.Marshal(bodyTranslate)
	if err != nil {
		log.Fatalf("json.Marshal[bodyTranslate]: %s", string(bodyTranslateBytes))
	}

	result, err := HttpRequest(http.MethodPost, TRANSLATE_API_URL, headers, bodyTranslateBytes)
	if err != nil {
		log.Fatalf("TranslateString [%s] %v", s, fmt.Errorf("HttpRequest: %w", err))
	}

	if result["translatedText"] == nil {
		log.Fatalf("TranslateString return a nil translatedText, query: %s, result: %v", s, result)
	}
	TOTAL_TRANSLATE += 1
	return result["translatedText"].(string), nil
}

func GetValue(item []byte) ([]byte, error) {
	var result []byte
	var rErr error
	var itemInterface interface{}
	err := json.Unmarshal(item, &itemInterface)
	if err != nil {
		return result, fmt.Errorf("GetValue.json.Unmarshal[item] %v", err)
	}
	valueOfItem := reflect.ValueOf(itemInterface)
	switch valueOfItem.Kind() {
	case reflect.Map:
		rM, err := GetMapValue(item)
		if err != nil {
			rErr = fmt.Errorf("GetMapValue.GetMapValue: %v", err)
		}
		result = rM
	case reflect.Slice:
		rS, err := GetSliceValue(item)
		if err != nil {
			rErr = fmt.Errorf("GetMapValue.GetSliceValue: %v", err)
		}
		result = rS
	case reflect.String:
		rT, err := TranslateString(strings.Trim(string(item), `"`))
		if err != nil {
			rErr = fmt.Errorf("TranslateString: %v", err)
		}
		//rT := strings.Trim(string(item), `"`) // test flow
		//result = []byte(fmt.Sprintf(`"...%s..."`, rT)) // test flow
		result = []byte(fmt.Sprintf(`"%s"`, rT))
	default:
		result = item
	}
	return result, rErr
}

func GetSliceValue(inputValue []byte) ([]byte, error) {
	var finalValue []byte
	var sliceResult [][]byte
	if err := json.Unmarshal(inputValue, &sliceResult); err != nil {
		return finalValue, fmt.Errorf("GetSliceValue.json.Unmarshal[inputValue] %v", err)
	}
	for _, item := range sliceResult {
		r, err := GetValue(item)
		if err != nil {
			return finalValue, fmt.Errorf("GetSliceValue.GetValue[item] %v", err)
		}
		sliceResult = append(sliceResult, r)
	}
	finalValue, err := json.Marshal(sliceResult)
	if err != nil {
		return finalValue, fmt.Errorf("GetSliceValue.json.Marshal[sliceResult] %v", err)
	}
	return finalValue, nil
}

func GetMapValue(jsonContent []byte) ([]byte, error) {
	var result []byte
	var finalValue []byte
	var gErr error
	var inputValue, resultValue orderedjson.Map
	if err := json.Unmarshal(jsonContent, &inputValue); err != nil {
		gErr = fmt.Errorf("GetMapValue.json.Unmarshal[jsonContent] %v", err)
		return finalValue, gErr
	}
	for _, mE := range inputValue {
		item := mE.Value
		var itemInterface interface{}
		err := json.Unmarshal(item, &itemInterface)
		if err != nil {
			return result, fmt.Errorf("GetMapValue.json.Unmarshal[item] %v", err)
		}
		result, err := GetValue(item)
		if err != nil {
			return finalValue, fmt.Errorf("GetMapValue.GetValue[item] %v", err)
		}
		entry := orderedjson.MapEntry{
			Key:   mE.Key,
			Value: json.RawMessage(result),
		}
		resultValue = append(resultValue, entry)
	}

	finalValue, err := json.MarshalIndent(resultValue, "", "\t")
	if err != nil {
		return finalValue, fmt.Errorf("GetMapValue.json.MarshalIndent[resultValue] %v", err)
	}
	return finalValue, gErr
}
