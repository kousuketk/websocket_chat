package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kousuketk/websocket_chat/server/model"
	"github.com/kousuketk/websocket_chat/server/service"
)

type MessageHandler struct {
	service service.MessageService
}

func NewMessageHandler() MessageHandler {
	return MessageHandler{
		service: service.NewMessageService(),
	}
}

func (m *MessageHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok."))
}

func (m *MessageHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket error: ", err)
		return
	}
	defer conn.Close()

	mt, wsbody, err := conn.ReadMessage()
	if err != nil {
		log.Println("Websocket error: ", err)
	}
	channelID := string(wsbody)

	for v := range m.service.Get(channelID) {
		res := v.(model.Message)
		j, err := json.Marshal(res)
		if err != nil {
			log.Println("Websocket error: ", err)
		} else {
			conn.WriteMessage(mt, j)
		}
	}
}

func (m *MessageHandler) Publish(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var msg model.Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = m.service.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
