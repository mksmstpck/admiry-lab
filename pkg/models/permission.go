package models

import "github.com/pborman/uuid"

type Permission struct {
	ID        uuid.UUID `json:"id" bun:"default:uuid_generate_v4() ,pk"`
	Name      string    `json:"name" binding:"required" bun:",unique"`
	AllowedTo []string  `json:"allowedTo" binding:"required" bun:",array"`
}
