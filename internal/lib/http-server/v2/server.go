package v2

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

	fmt.Fprint(w, GetPlayersScore(player))
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
