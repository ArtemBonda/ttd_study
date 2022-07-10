package v1

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayServer(w http.ResponseWriter, r *http.Request) {
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
