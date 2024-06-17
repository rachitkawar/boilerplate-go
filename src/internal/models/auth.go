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
	UserId          int    `json:"user_id"`
	RoleId          int    `json:"role_id"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Verified        bool   `json:"verified"`
	ProfileComplete bool   `json:"profile_complete"`
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

type GoogleLoginRequest struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

type SocialLoginModel struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
