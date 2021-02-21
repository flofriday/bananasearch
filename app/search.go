package app

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func (s *Server) handleSearch() http.HandlerFunc {
	landingPage, _ := ioutil.ReadFile("web/index.html")
	resultPage, _ := ioutil.ReadFile("web/results.html")
	resultTemplate, _ := template.New("result").Parse(string(resultPage))

	type resultModel struct {
		Query    string
		Duration time.Duration
		Results  []string
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		searchTerm := r.URL.Query().Get("q")

		// Check if we have a search term
		if searchTerm == "" {
			rw.WriteHeader(http.StatusOK)
			rw.Write(landingPage)
			return
		}
		searchTerm = strings.ToLower(searchTerm)

		start := time.Now()
		results := s.Index.GetDocs(searchTerm)
		duration := time.Since(start)

		result := resultModel{
			Query:    searchTerm,
			Duration: duration,
			Results:  results,
		}
		err := resultTemplate.Execute(rw, result)
		if err != nil {
			log.Printf("Template ERROR: %s", err.Error())
		}
	}
}
