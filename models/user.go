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
	Id               uint `gorm:"primaryKey"`
	FullName         string
	UserName         string
	Password         string
	Email            string `json:"username" gorm:"unique"`
	Role             Role
	VerificationCode string
	Verified         bool `gorm:"not null"`
}

type SignUpInput struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
}

type SignInInput struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID        uint
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
