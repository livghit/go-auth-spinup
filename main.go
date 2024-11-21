package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name":     "pedro",
				"email":    "pedro",
				"password": "salted-password",
			},
		}

		writeJSON(data, w)
	})

	log.Default().Printf("Server started")

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func writeJSON(data map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
