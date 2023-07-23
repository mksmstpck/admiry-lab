package models

type AuthByUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthByEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
