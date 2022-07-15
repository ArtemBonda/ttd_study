package v1

//func TestGETPlayer(t *testing.T) {
//	t.Run("returns Pepper's score", func(t *testing.T) {
//		request, err := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		response := httptest.NewRecorder()
//
//		PlayServer(response, request)
//		got := response.Body.String()
//		want := "20"
//
//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//	})
//	t.Run("returns Floyd's score", func(t *testing.T) {
//		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
//		response := httptest.NewRecorder()
//
//		PlayServer(response, request)
//
//		got := response.Body.String()
//		want := "10"
//
//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//	})
//}
