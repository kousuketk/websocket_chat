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


### look redis
```
$ docker exec -it redis-server bash
root@:/# redis-cli
```