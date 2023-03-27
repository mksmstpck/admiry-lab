package models

type Response struct {
	Status  int32       `json:"status"`
	Message interface{} `json:"message"`
}
