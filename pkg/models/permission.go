package models

import "github.com/pborman/uuid"

type Permission struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" binding:"required"`
	AllowedTo []string  `json:"allowedTo" binding:"required"`
}
