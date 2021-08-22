package redisRepository

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/kousuketk/websocket_chat/server/model"
)

type RedisMessageRepository struct{}

func NewRedisMessageRepository() RedisMessageRepository {
	return RedisMessageRepository{}
}

func (r *RedisMessageRepository) GetMessage(channelID string) error {
	conn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channelID)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			log.Printf("%s: message : %s\n", v.Channel, v.Data)
		case redis.Subscription:
			log.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return error
		}
	}
}

func (r *RedisMessageRepository) SendMessage(msg model.Message) error {
	conn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r, err2 := conn.Do("PUBLISH", msg.GetChannelID(), msg.GetContent())
	if err2 != nil {
		panic(err2)
	}
	log.Println(r)

	return nil
}
