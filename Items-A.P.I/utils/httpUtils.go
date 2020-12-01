package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON : To give the repsonse back to the client in json form
func RespondJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
