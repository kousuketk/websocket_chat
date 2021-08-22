package repository

import "github.com/kousuketk/websocket_chat/server/model"

type MessageRepository interface {
	GetMessage(channelID string) error
	SendMessage(msg model.Message) error
}
