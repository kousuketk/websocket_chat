# websocket_chat

### start
```
$ docker-compose up -d
```

### run server
```
$ docker exec -it websocket_chat bash
$ go run app/server/main.go 
```

### PubSub(Publish via golang, Subscribe via wscat)
```
- Publish
$ curl -X POST http://localhost:9090/api/v1/publish -d '{"channelID":"1","userID":"1","content":"test-con","sentAt":"test-sentAt"}'

- Subscribe
$ wscat -c ws://0.0.0.0:9090/api/v1/subscribe
Connected (press CTRL+C to quit)
> 1
< {"channelID":"1","userID":"1","content":"test-content1","sentAt":"test-sentAt"}
```

### look redis
```
$ docker exec -it redis-server bash
root@:/# redis-cli
```