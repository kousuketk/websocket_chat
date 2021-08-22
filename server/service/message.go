package service

import (
	"github.com/kousuketk/websocket_chat/server/model"
	"github.com/kousuketk/websocket_chat/server/registry"
	"github.com/kousuketk/websocket_chat/server/repository"
)

type MessageService struct {
	repo repository.MessageRepository
}

func NewMessageService() MessageService {
	return MessageService{
		repo: registry.NewRedisMessageRepository(),
	}
}

func (m *MessageService) Send(msg model.Message) error {
	err := m.repo.SendMessage(msg)
	if err != nil {
		return nil
	}
	return nil
}

func (m *MessageService) Get(channelID string) error {
	err := m.repo.GetMessage(channelID)
	if err != nil {
		return nil
	}
	return nil
}
