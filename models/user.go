package models

import "gorm.io/gorm"

type Role int

const (
	USER Role = iota
	ADMIN
	SUPER_ADMIN
)

type User struct {
	gorm.Model
	Id       uint `gorm:"primaryKey"`
	FullName string
	UserName string
	Password string
	Email    string `json:"username" gorm:"unique"`
	Role     Role
}
