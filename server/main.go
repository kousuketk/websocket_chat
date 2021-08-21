package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/kousuketk/websocket_chat/server/controller"
)

func setRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		m := controller.MessageController{}
		r.Get("/health", m.Health)
		r.Get("/subscribe", m.Subscribe)
		r.Post("/publish", m.Publish)
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
