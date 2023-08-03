package server

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func Start() {
	muxHttp := http.NewServeMux()
	fileServers := http.FileServer(http.Dir(PUBLIC_PATH))
	muxHttp.Handle("/", fileServers)
	muxHttp.HandleFunc("/home/", getIndexHandler)
	muxHttp.HandleFunc("/hero/", getHeroListHandler)
	muxHttp.HandleFunc("/hero-detail/", getHeroDetailHandler)
	muxHttp.HandleFunc("/chat/", chatFormHandler)

	muxHttp.HandleFunc("/base", getBaseHandler)
	muxHttp.HandleFunc("/another-path", getAnotherHandler)
	muxHttp.HandleFunc("/json-response", getJsonResponseHandler)
	muxHttp.HandleFunc("/request-timeout", getRequestTimeoutHandler)
	muxHttp.HandleFunc("/post-request", postRequestHandler)
	muxHttp.Handle("/chat-messages/", websocket.Handler(websocketChatHandler))

	fmt.Println("All path registered, server is run on port 8080")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      muxHttp,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		serverTLS := &http.Server{
			Addr:         ":8081",
			Handler:      muxHttp,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}
		err := serverTLS.ListenAndServeTLS("course/http/server/tls/crt.pem", "course/http/server/tls/key.pem")
		if err != nil {
			fmt.Println("serverTLS.ListenAndServeTLS err", err)
		}
	}()

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server.ListenAndServe err", err)
	}

}
