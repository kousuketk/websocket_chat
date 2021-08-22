package handler

import (
	"encoding/json"
	"fmt"
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
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	_, b, err2 := conn.ReadMessage()
	if err2 != nil {
		log.Println("Websocket error: ", err2)
	}
	channelID := string(b)

	err3 := m.service.Get(channelID)
	if err3 != nil {
		log.Println(err3)
	}

	// for v := range m.service.Get(channelID) {
	// 	switch v.(type) {
	// 	case model.Message:
	// 		res := v.(model.Message)
	// 		j, err := json.Marshal(res)
	// 		if err != nil {
	// 			log.Println("Websocket error: ", err)
	// 		} else {
	// 			conn.WriteMessage(mt, j)
	// 		}
	// 	}
	// }

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

	err2 := m.service.Send(msg)
	if err2 != nil {
		log.Println(err2)
	}
	return
}
