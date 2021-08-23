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

func (m *MessageService) Get(channelID string) chan interface{} {
	ch := m.repo.GetMessage(channelID) // redisRepoからのchanelを受け取ってそれをただhandlerに返すだけでもいい感じしてきた
	return ch
	// // channelを返す
	// ch := make(chan interface{})

	// go func() {
	// 	for v := range m.repo.GetMessage(channelID) {
	// 		switch v.(type) {
	// 		case model.Message:
	// 			log.Println("catched in service")
	// 			ch <- ch
	// 		case error:
	// 			log.Println("servecei error")
	// 		}
	// 	}
	// 	// close(ch)
	// }()
	// return ch
}
