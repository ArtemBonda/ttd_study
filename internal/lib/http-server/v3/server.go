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
	players := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		s.ShowScore(w, players)
	case http.MethodPost:
		s.processWin(w, players)
	}
}

func (s *PlayerServer) ShowScore(w http.ResponseWriter, playerName string) {
	score := s.Store.GetPlayerScore(playerName)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, playerName string) {
	s.Store.RecordWin(playerName)
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
