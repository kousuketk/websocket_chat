package service

import (
	"github.com/kousuketk/websocket_chat/server/model"
	"github.com/kousuketk/websocket_chat/server/repository"
)

type MessageService struct{}

var messageRepo repository.MessageRepository

func (MessageService) Send(msg model.Message) error {
	err := messageRepo.SendMessage(msg)
	if err != nil {
		return nil
	}
	return nil
}

func (MessageService) Get(channelID string) error {
	err := messageRepo.GetMessage(channelID)
	if err != nil {
		return nil
	}
	return nil
}
