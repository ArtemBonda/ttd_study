package v3

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) string
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	Store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.ShowScore(w, r)
	case http.MethodPost:
		s.processWin(w, r)
	}
}

func (s *PlayerServer) ShowScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := s.Store.GetPlayerScore(player)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	s.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}
	return ""
}
