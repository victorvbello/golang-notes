package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

const (
	PUBLIC_PATH = "./course/http/server/public/"
)

var publicHtmlTemplates *template.Template

func init() {
	var err error
	publicHtmlTemplates, err = template.ParseFiles(
		PUBLIC_PATH+"/menu.html",
		PUBLIC_PATH+"/index.html",
		PUBLIC_PATH+"hero/list.html",
		PUBLIC_PATH+"/hero/detail.html",
		PUBLIC_PATH+"/chat.html",
	)
	if err != nil {
		fmt.Println("error on load template.ParseFiles", err)
		return
	}
}

func getBaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi from baseHandler")
}

func getAnotherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi from anotherHandler")
}

func getJsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, `{"id":1, "name": "test"}`)
}

func postRequestHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody interface{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println("postHandler decore body err", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":decore body err}`)
		return
	}
	b, err := json.Marshal(struct {
		Payload interface{}
	}{
		Payload: requestBody,
	})
	if err != nil {
		fmt.Println("postHandler json Marshal err", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":marshal json response err}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(b)) //  json.NewEncoder(w).Encode(responseBody)
}

func getRequestTimeoutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprint(w, "hi from getRequestTimeoutHandler")
}

func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	err := publicHtmlTemplates.ExecuteTemplate(w, "index.html", DB.Heroes)
	if err != nil {
		fmt.Println("getIndexHandler publicHtmlTemplates.ExecuteTemplate", err)
		return
	}
}

func getHeroListHandler(w http.ResponseWriter, r *http.Request) {
	err := publicHtmlTemplates.ExecuteTemplate(w, "list.html", DB.Heroes)
	if err != nil {
		fmt.Println("getHeroListHandler publicHtmlTemplates.ExecuteTemplate", err)
		return
	}
}

func getHeroDetailHandler(w http.ResponseWriter, r *http.Request) {
	inputId := strings.TrimPrefix(r.URL.Path, "/hero-detail/")
	id, err := strconv.Atoi(inputId)
	if err != nil {
		fmt.Println("strconv.Atoi", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":invalid id}`)
		return
	}

	var foundHero *Hero

	for _, hero := range DB.Heroes {
		if hero.ID == id {
			foundHero = &hero
			break
		}
	}

	if foundHero == nil {
		fmt.Println("strconv.Atoi", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fmt.Sprintf(`{"info":hero id:%d not found}`, id))
	}

	err = publicHtmlTemplates.ExecuteTemplate(w, "detail.html", foundHero)
	if err != nil {
		fmt.Println("getHeroDetailHandler publicHtmlTemplates.ExecuteTemplate", err)
		return
	}
}

func chatFormHandler(w http.ResponseWriter, r *http.Request) {
	err := publicHtmlTemplates.ExecuteTemplate(w, "chat.html", DB.Heroes)
	if err != nil {
		fmt.Println("chatFormHandler publicHtmlTemplates.ExecuteTemplate", err)
		return
	}
}

func websocketChatHandler(wsCnn *websocket.Conn) {
	name := wsCnn.Request().URL.Query().Get("name")
	CHAT_CLIENTS[int(time.Now().UnixNano())] = wsCnn
	go func() {
		for msg := range CHAT_MSG {
			for cId, cWsCnn := range CHAT_CLIENTS {
				err := websocket.Message.Send(cWsCnn, msg)
				if err != nil {
					fmt.Println("websocketChatHandler websocket.Message.Send", err)
					delete(CHAT_CLIENTS, cId)
					continue
				}
			}
		}
	}()

	for {
		var msg string
		err := websocket.Message.Receive(wsCnn, &msg)
		if err != nil {
			fmt.Println("websocketChatHandler websocket.Message.Receive", err)
			return
		}
		CHAT_MSG <- fmt.Sprintf("<%s>: %s", name, msg)
	}
}
