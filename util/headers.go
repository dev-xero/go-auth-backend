package util

import "net/http"

func SetJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
