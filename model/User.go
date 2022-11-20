package model

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// User is a struct that represents a user
type User struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	Username string
	Password string
	Fullname string
	Avatar   string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
