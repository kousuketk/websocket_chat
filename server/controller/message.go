package controller

import (
	"fmt"
	"net/http"

	"github.com/kousuketk/websocket_chat/server/service"
)

type MessageController struct{}

var messageService service.MessageService

func (m *MessageController) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok."))
}

func (m *MessageController) Subscribe(w http.ResponseWriter, r *http.Request) {
	// upgrader := websocket.Upgrader{
	// 	ReadBufferSize:  1024,
	// 	WriteBufferSize: 1024,
	// 	CheckOrigin: func(r *http.Request) bool {
	// 		return true
	// 	},
	// }
	// conn, err := upgrader.Upgrade(w, r, nil)
	// defer conn.Close()
	w.Write([]byte(fmt.Sprintln("ok. sub")))
}

func (m *MessageController) Publish(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintln("ok. pub")))
}
