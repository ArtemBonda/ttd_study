package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayer(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Papper", nil)
		response := httptest.NewRecorder()

		PlayServer(response, request)
		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("return Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayServer(response, request)

		got := response.Body.String()
		want := "30"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
