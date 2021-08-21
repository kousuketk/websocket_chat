package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/kousuketk/websocket_chat/server/controller"
)

func setRouter() http.Handler {
	r := chi.newRouter()
	r.Route("api/v1", func(r chi.Router) {
		m := controller.MessageController{}
		r.Get("/subscribe", m.Subscribe)
		r.Post("/publish", m.Publish)
	})
	return r
}

func main() {
	r := setRouter()
	http.ListenAndServe(os.Getenv("API_PORT"), r)
}
