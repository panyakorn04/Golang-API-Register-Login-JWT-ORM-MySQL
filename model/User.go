package model

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// User is a struct that represents a user
type User struct {
	gorm.Model
	UserID          uint   `json:"user_id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	Picture_profile string `json:"picture_profile"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
