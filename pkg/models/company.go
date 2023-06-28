package models

import "github.com/pborman/uuid"

type Company struct {
	ID      uuid.UUID   `json:"id" bun:"default:uuid_generate_v4() ,pk"`
	Name    string      `json:"name" validate:"required" bun:",unique"`
	UserIDs []uuid.UUID `json:"user_ids" bun:",array"`
}
