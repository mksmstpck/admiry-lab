package models

type Permission struct {
	ID        int32    `json:"id"`
	Name      string   `json:"name" binding:"required"`
	AllowedTo []string `json:"allowedTo" binding:"required"`
}
