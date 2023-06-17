package models

type Role struct {
	ID           int32  `json:"id"`
	Name         string `json:"name" validate:"required"`
	PermissionID string `json:"permissions" validate:"required"`
}
