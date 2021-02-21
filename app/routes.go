package app

import (
	"github.com/go-chi/chi/middleware"
)

func (s *Server) Routes() {
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", s.handleSearch())
}
