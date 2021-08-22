package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kousuketk/websocket_chat/server/model"
	"github.com/kousuketk/websocket_chat/server/registry"
	"github.com/kousuketk/websocket_chat/server/repository"
	"github.com/kousuketk/websocket_chat/server/service"
)

type MessageHandler struct {
	repo repository.MessageRepository
}

func NewMessageHandler() MessageHandler {
	return MessageHandler{
		repo: registry.NewRedisMessageRepository(),
	}
}

var messageService service.MessageService

func (m *MessageHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok."))
}

func (m *MessageHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	mt, b, err := conn.ReadMessage()
	if err != nil {
		log.Println("Websocket error: ", err)
	}
	channelID := string(b)

	for v := range messageService.Get(channelID) {
		switch v.(type) {
		case model.Message:
			res := v.(model.Message)
			j, err := json.Marshal(res)
			if err != nil {
				log.Println("Websocket error: ", err)
			} else {
				conn.WriteMessage(mt, j)
			}
		}
	}

	w.Write([]byte(fmt.Sprintln("ok. sub")))
}

func (m *MessageHandler) Publish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var msg model.Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err2 := messageService.Send(msg)
	if err2 != nil {
		log.Println(err2)
	}
	return
}
