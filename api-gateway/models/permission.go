package models

type Permission struct {
	ID        int32    `json:"id"`
	Name      string   `json:"name" validate:"required"`
	AllowedTo []string `json:"allowedTo" validate:"required"`
}
