package models

type Response[T any] struct {
	Status  int32  `json:"status"`
	Message T      `json:"message"`
	Error   string `json:"error"`
}
