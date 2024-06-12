package models

import "github.com/dgrijalva/jwt-go"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Expiry string `json:"expiry"`
}

type Claims struct {
	UserId    int    `json:"user_id"`
	RoleId    int    `json:"role_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.StandardClaims
}
