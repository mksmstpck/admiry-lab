package models

import "github.com/pborman/uuid"

type User struct {
	ID       uuid.UUID `json:"id" bun:"default:uuid_generate_v4() ,pk"`
	Username string    `json:"username" validate:"required" bun:",unique"`
	FullName string    `json:"fullName" validate:"required" bun:",unique"`
	Email    string    `json:"email" validate:"required,email" bun:",unique"`
	Password string    `json:"password" validate:"required" bun:",notnull"`
}
