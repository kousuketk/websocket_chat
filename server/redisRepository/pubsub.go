package redisRepository

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/kousuketk/websocket_chat/server/model"
)

type RedisMessageRepository struct{}

func NewRedisMessageRepository() RedisMessageRepository {
	return RedisMessageRepository{}
}

func (repo RedisMessageRepository) GetMessage(channelID string) chan interface{} { // pointer recieverにしたらエラーでた
	conn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channelID)
	ch := make(chan interface{})
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				log.Printf("%s: message : %s\n", v.Channel, v.Data)
				var msg model.Message
				err := json.Unmarshal(v.Data, &msg)
				if err != nil {
					log.Println("pubsub, unmarshal error")
				}
				ch <- msg
			case redis.Subscription:
				log.Printf("channelID:%s: Kind:%s Count:%d\n", v.Channel, v.Kind, v.Count)
			case error:
				ch <- v
			}
		}
		// close(ch)
	}()
	return ch
}

func (repo RedisMessageRepository) SendMessage(msg model.Message) error {
	conn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	j, err3 := json.Marshal(msg)
	if err3 != nil {
		panic(err3)
	}

	_, err2 := conn.Do("PUBLISH", msg.GetChannelID(), string(j))
	if err2 != nil {
		panic(err2)
	}

	return nil
}
