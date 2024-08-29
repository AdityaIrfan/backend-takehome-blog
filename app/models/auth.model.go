package models

type Register struct {
	Name     string `json:"Name" form:"Name" validate:"required"`
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type Login struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"RefreshToken" form:"RefreshToken" validate:"required"`
}

type TokenResponse struct {
	Token string `json:"Token"`
}
