package registry

import (
	"github.com/kousuketk/websocket_chat/server/redisRepository"
	"github.com/kousuketk/websocket_chat/server/repository"
)

func NewRedisMessageRepository() repository.MessageRepository {
	return redisRepository.NewRedisMessageRepository()
}
