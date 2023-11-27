package main

import (
	"go-practice/middleware/md"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// この書き方をするとOne,Towの順番で実行される
		r.Use(md.MiddlewareOne, md.MiddlewareTwo)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello new world!"))
		})
	})
	http.ListenAndServe(":8080", r)
}
