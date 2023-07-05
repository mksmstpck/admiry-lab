package models

import "github.com/pborman/uuid"

type User struct {
	ID       uuid.UUID `json:"id" bun:"default:uuid_generate_v1() ,pk"`
	Username string    `json:"username" binding:"required" bun:",unique"`
	FullName string    `json:"fullName" binding:"required" bun:",unique"`
	Email    string    `json:"email" binding:"required" bun:",unique"`
	Password string    `json:"password" binding:"required" bun:",notnull"`
}
