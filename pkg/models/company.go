package models

import "github.com/pborman/uuid"

type Company struct {
	ID      uuid.UUID   `json:"id" bun:"default:uuid_generate_v1() ,pk"`
	Name    string      `json:"name" binding:"required" bun:",unique"`
	UserIDs []uuid.UUID `json:"user_ids" bun:",array"`
}
