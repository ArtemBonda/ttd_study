package v3

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) string
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	Store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	score := s.Store.GetPlayerScore(player)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func GetPlayersScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}
	return ""
}
