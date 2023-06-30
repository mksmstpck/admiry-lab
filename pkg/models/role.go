package models

import "github.com/pborman/uuid"

type Role struct {
	ID            uuid.UUID   `json:"id" bun:"default:uuid_generate_v4() ,pk"`
	Name          string      `json:"name" validate:"required"`
	PermissionIDs []uuid.UUID `json:"permissions" bun:",array"`
}
