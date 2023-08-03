package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func processResponse(r *http.Response) error {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll %v", err)
	}
	fmt.Println("processResponse => ", r.Request.Method, r.StatusCode, string(b))
	return nil
}

func Start() {
	hostURL := "http://127.0.0.1:8080"

	clientHttp := http.Client{
		Timeout: 2 * time.Second,
	}

	//Get
	fmt.Println("client is start")
	getPathS := []string{"/base", "/another-path", "/json-response", "/request-timeout"}
	for _, path := range getPathS {
		fmt.Println("-- GET", path)
		resp, err := clientHttp.Get(hostURL + path)
		if err != nil {
			fmt.Println("error http.Get", err)
			continue
		}
		processResponse(resp)
	}

	fmt.Println("-- POST /")

	b, err := json.Marshal(struct {
		ID   int
		Name string
	}{
		ID:   123,
		Name: "test-client",
	})
	if err != nil {
		fmt.Println("json.Marshal post", err)
		return
	}

	resp, err := clientHttp.Post(hostURL+"/post-request", "application/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println("error http.Post", err)
		return
	}
	processResponse(resp)
}
