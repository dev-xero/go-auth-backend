package util

type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}
