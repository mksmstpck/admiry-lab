package models

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	RoleID   int    `json:"role" validate:"required"`
}
