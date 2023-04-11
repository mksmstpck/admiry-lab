package models

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username" validate:"required,name"`
	FullName string `json:"fullName" validate:"required,fullname"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	RoleID   int    `json:"role" validate:"required,role"`
}
