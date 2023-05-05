package models

import (
	"time"

	"gorm.io/gorm"
)

type Role int

const (
	USER Role = iota
	ADMIN
	SUPER_ADMIN
)

type User struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey"`
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Role     Role
}

type SignUpInput struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id        uint
	FullName  string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
