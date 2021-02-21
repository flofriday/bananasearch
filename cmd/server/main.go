package main

import (
	"log"
	"net/http"

	"github.com/flofriday/websearch/app"
	"github.com/flofriday/websearch/store"
	"github.com/go-chi/chi"
)

func main() {
	index, err := store.LoadIndex("index")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Websites in Index: %d", index.NumDocs())
	log.Printf("Words in Index: %d", index.NumWords())
	server := app.Server{
		Index:  index,
		Router: chi.NewRouter(),
	}
	server.Routes()

	http.ListenAndServe("localhost:8000", &server)
}
