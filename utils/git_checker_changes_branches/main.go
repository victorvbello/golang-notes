package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type requestResult map[string]interface{}
type requestHeaders map[string]string
type repoInputItem struct {
	Repo       string `json:"repo"`
	Branch1    string `json:"branch1"`
	Branch2    string `json:"branch2"`
	BrowserUrl string `json:"browser_url"`
}

func HttpRequest(method string, url string, headers requestHeaders) (requestResult, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	var resultBody requestResult
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return resultBody, err
	}
	if len(headers) > 0 {
		for key, data := range headers {
			req.Header.Set(key, data)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return resultBody, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resultBody, err
	}
	err = json.Unmarshal(body, &resultBody)
	if err != nil {
		return resultBody, err
	}
	return resultBody, nil
}

func GetRepoInputItems(pathFile string) ([]repoInputItem, error) {
	var result []repoInputItem
	buffer, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return result, fmt.Errorf("ioutil.ReadFile: %w", err)
	}
	if err = json.Unmarshal(buffer, &result); err != nil {
		return result, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return result, nil
}

func main() {
	args := os.Args[1:]
	var headers requestHeaders
	var token string
	var fileInputPath string
	var err error
	var readyToProd []repoInputItem

	if len(os.Args) < 2 {
		log.Fatal("token and fileInputPath is required")
	}

	token = args[0]
	fileInputPath = args[1]

	fmt.Println("token", token)
	fmt.Println("fileInputPath", fileInputPath)

	headers = requestHeaders(map[string]string{
		"Authorization": "Bearer " + token,
	})

	repos, err := GetRepoInputItems(fileInputPath)
	if err != nil {
		log.Fatal(fmt.Errorf("GetRepoInputItems: %w", err))
	}

	for _, r := range repos {
		result, err := HttpRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/compare/%s...%s", r.Repo, r.Branch1, r.Branch2), headers)
		if err != nil {
			log.Fatal(fmt.Errorf("HttpRequest: %w", err))
		}
		total_commits := result["total_commits"].(float64)
		fmt.Printf("Process [%s]\n", r.Repo)
		if total_commits > 0 {
			r.BrowserUrl = fmt.Sprintf("https://github.com/%s/compare/%s...%s", r.Repo, r.Branch1, r.Branch2)
			readyToProd = append(readyToProd, r)
		}
	}
	fmt.Println("----")
	if len(readyToProd) > 0 {
		fmt.Println("Repos ready to Prod")
	} else {
		fmt.Println("Has no repo to Prod")
	}
	for _, r := range readyToProd {
		fmt.Printf("[%s] %s\n", r.Repo, r.BrowserUrl)
	}
	fmt.Println("----")
}
