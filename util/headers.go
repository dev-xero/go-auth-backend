package util

import "net/http"

/*
Sets the response headers content type to application/json

Objectives:
  - Set the response content type header to a json type

Params:
  - w: A http response writer

Returns:
  - No return value
*/
func SetJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
