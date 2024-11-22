package handlers

import (
	"io"
	"net/http"

	"github.com/livghit/go-auth-spinup/internal/api"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		io.WriteString(w, "404 Not Found")
		return
	}
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
}

func Hi(w http.ResponseWriter, _ *http.Request) {
	api.WriteJSON(map[string]interface{}{"hi": "man"}, w)
}
