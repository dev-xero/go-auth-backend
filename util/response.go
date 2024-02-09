package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

/*
Response body struct

Fields:
  - Message: string
  - Success: bool
  - Payload: interface{} (any valid go data type)
*/
type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

/*
User payload struct

Fields:
  - ID:       uuid
  - Username: string
  - Email:    string
*/
type UserPayload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

/*
Sends a JSON response to the client with an optional payload

Objectives:
  - Set the response content headers to a json type
  - Build the response object with the message and payload
  - Encode the response object as a json object
  - Write the header and response body with the json object

Params:
  - w:       A http response writer
  - msg:     The response body message
  - status:  The response status
  - payload: An optional payload

Returns:
  - No return value
*/
func JsonResponse(w http.ResponseWriter, msg string, status int, payload interface{}) {
	SetJSONHeaders(w)

	res := Response{
		Message: msg,
		Success: status < 400,
		Payload: payload,
	}

	// Encode the response object
	json, err := json.Marshal(res)
	if err != nil {
		errMsg := fmt.Errorf("[FAIL]: failed to encode json response: %w", err)
		http.Error(w, errMsg.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(json)
}
