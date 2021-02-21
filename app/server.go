package app

import (
	"net/http"

	"github.com/flofriday/websearch/store"
	"github.com/go-chi/chi"
)

type Server struct {
	Index  *store.Index
	Router *chi.Mux
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(rw, r)
}
