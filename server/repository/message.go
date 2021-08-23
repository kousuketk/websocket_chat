package repository

import "github.com/kousuketk/websocket_chat/server/model"

type MessageRepository interface {
	GetMessage(channelID string) chan interface{}
	SendMessage(msg model.Message) error
}
