package dto

type APIResponse[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
