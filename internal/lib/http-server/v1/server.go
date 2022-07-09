package v1

import (
	"fmt"
	"net/http"
)

func PlayServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
