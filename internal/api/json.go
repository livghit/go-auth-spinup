package api

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(data map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
