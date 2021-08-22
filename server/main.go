package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/kousuketk/websocket_chat/server/handler"
)

func setRouter() http.Handler {
	r := chi.NewRouter()
	h := handler.NewMessageHandler()
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", h.Health)
		r.Get("/subscribe", h.Subscribe)
		r.Post("/publish", h.Publish)
	})
	return r
}

func main() {
	// 環境変数の設定
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	r := setRouter()
	http.ListenAndServe(":"+os.Getenv("API_PORT"), r)
}
