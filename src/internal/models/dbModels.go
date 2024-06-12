package models

import "time"

// --- Define all the table structures here -- //

// UserDb <<--User Model-->>
type UserDb struct {
	Id          int       `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Password    string    `json:"password" db:"password"`
	Email       string    `json:"email" db:"email"`
	RoleId      int       `json:"role_id" db:"role_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// RolesDb <<--Role Model-->>
type RolesDb struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// PermissionsDb <<--Permission Model-->>
type PermissionsDb struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	RoleId      int       `json:"role_id" db:"role_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
