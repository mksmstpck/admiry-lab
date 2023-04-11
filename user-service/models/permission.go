package models

type Permission struct {
	ID        int32  `json:"id"`
	Name      string `json:"name" validate:"required,name"`
	AllowedTo string `json:"allowedTo" validate:"required,allowedTo"`
}
