package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

type UserPayload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func JsonResponse(w http.ResponseWriter, msg string, status int, payload interface{}) {
	SetJSONHeaders(w)

	res := Response{
		Message: msg,
		Success: status < 400,
		Payload: payload,
	}

	json, err := json.Marshal(res)
	if err != nil {
		errMsg := fmt.Errorf("[FAIL]: failed to encode json response: %w", err)
		http.Error(w, errMsg.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(json)
}
