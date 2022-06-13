package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	GITHUB_MAIN_BRANCHES    = "main,master"
	GITHUB_DEVELOP_BRANCHES = "develop"
	GITHUB_API_BASE_URL     = "https://api.github.com"
	GITHUB_WEB_URL          = "https://github.com"
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resultBody, fmt.Errorf("ioutil.ReadAll %w", err)
	}
	err = json.Unmarshal(body, &resultBody)
	if err != nil {
		return resultBody, fmt.Errorf("json.Unmarshal %w", err)
	}
	return resultBody, nil
}

func HttpRequestSlice(method string, url string, headers requestHeaders) ([]requestResult, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	var resultBody []requestResult
	req, err := http.NewRequest(method, url, nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resultBody, fmt.Errorf("ioutil.ReadAll %w", err)
	}
	err = json.Unmarshal(body, &resultBody)
	if err != nil {
		return resultBody, fmt.Errorf("json.Unmarshal %w", err)
	}
	return resultBody, nil
}

func main() {
	args := os.Args[1:]
	var token string
	var organization string
	var readyToProd []repoInputItem
	var headers requestHeaders

	if len(os.Args) < 2 {
		log.Fatal("token and fileInputPath is required")
	}

	organization = args[0]
	token = args[1]

	log.Println("token", token)

	headers = requestHeaders(map[string]string{
		"Authorization": "Bearer " + token,
	})

	pageRepos := 1
	allRepos := []requestResult{}

	for {
		log.Println("Found repos")
		currentRepos, err := HttpRequestSlice("GET", fmt.Sprintf("%s/orgs/%s/repos?per_page=100&type=private&sort=updated&page=%d", GITHUB_API_BASE_URL, organization, pageRepos), headers)
		if err != nil {
			log.Fatal(fmt.Errorf("HttpRequestSlice: %w", err))
			break
		}
		if len(currentRepos) == 0 {
			log.Println("Total repos found:", len(allRepos))
			break
		}
		pageRepos += 1
		allRepos = append(allRepos, currentRepos...)
	}

	var wg sync.WaitGroup
	repoChan := make(chan repoInputItem)
	wg.Add(len(allRepos))
	for rI, r := range allRepos {
		repoUpdate := r["updated_at"].(string)
		uTime, err := time.Parse(time.RFC3339, repoUpdate)

		if err != nil {
			log.Fatalf("[%d] %v", rI, fmt.Errorf(" time.Parse: %w", err))
			continue
		}
		if uTime.Year() < time.Now().Year() {
			wg.Done()
			continue
		}
		repoName := r["full_name"].(string)
		go func(cwg *sync.WaitGroup, rName string, index int) {
			defer func() {
				cwg.Done()
			}()
			cRepo := repoInputItem{
				Repo: rName,
			}
			resultBranches, err := HttpRequestSlice("GET", fmt.Sprintf("%s/repos/%s/branches?protected=true", GITHUB_API_BASE_URL, rName), headers)
			if err != nil {
				log.Fatalf("[%d] %v", index, fmt.Errorf("HttpRequest: %w", err))
				return
			}

			if len(resultBranches) == 0 {
				resultBranches, err = HttpRequestSlice("GET", fmt.Sprintf("%s/repos/%s/branches", GITHUB_API_BASE_URL, rName), headers)
				if err != nil {
					log.Fatalf("[%d] %v", index, fmt.Errorf("HttpRequest: %w", err))
					return
				}
			}

			for _, r := range resultBranches {
				branch := r["name"].(string)
				if bytes.Contains([]byte(GITHUB_MAIN_BRANCHES), []byte(branch)) {
					cRepo.Branch1 = branch
				}
				if bytes.Contains([]byte(GITHUB_DEVELOP_BRANCHES), []byte(branch)) {
					cRepo.Branch2 = branch
				}
			}

			if cRepo.Branch1 == "" {
				return
			}

			if cRepo.Branch2 == "" {
				cRepo.Branch2 = cRepo.Branch1
			}

			log.Printf("[%d] Opted [%s] (%s) (%s)\n", index, cRepo.Repo, cRepo.Branch1, cRepo.Branch2)

			result, err := HttpRequest("GET", fmt.Sprintf("%s/repos/%s/compare/%s...%s", GITHUB_API_BASE_URL, cRepo.Repo, cRepo.Branch1, cRepo.Branch2), headers)
			if err != nil {
				log.Fatalf("[%d] %v", index, fmt.Errorf("HttpRequest: %w", err))
			}
			if result["total_commits"] == nil {
				log.Fatalf("[%d] %v", index, fmt.Errorf("total_commits, is nil"))
			}
			total_commits := result["total_commits"].(float64)
			log.Printf("[%d] Process [%s]\n", index, cRepo.Repo)
			if total_commits > 0 {
				cRepo.BrowserUrl = fmt.Sprintf("%s/%s/compare/%s...%s", GITHUB_WEB_URL, cRepo.Repo, cRepo.Branch1, cRepo.Branch2)
				repoChan <- cRepo
			}

		}(&wg, repoName, rI)
	}

	go func() {
		wg.Wait()
		close(repoChan)
	}()

	for r := range repoChan {
		readyToProd = append(readyToProd, r)
	}

	if len(readyToProd) > 0 {
		fmt.Println("Repos ready to Prod")
	} else {
		fmt.Println("Has no repo to Prod")
	}

	for _, r := range readyToProd {
		fmt.Printf("[%s] %s\n", r.Repo, r.BrowserUrl)
	}
}
