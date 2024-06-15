package models

import "github.com/dgrijalva/jwt-go"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	UserId    int    `json:"user_id"`
	RoleId    int    `json:"role_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Verified  bool   `json:"verified"`
	jwt.StandardClaims
}

type SignupRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding: "required"`
}

type SignupResponse struct {
	Token string `json:"token"`
}

type VerifyRequest struct {
	Token string `json:"token" binding:"required"`
}

type VerifyResponse struct {
	UserId    int    `json:"user_id"`
	RoleId    int    `json:"role_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Verified  bool   `json:"verified"`
}
