package database

import "github.com/rachitkawar/boilerplate-go/src/internal/models"

type Store interface {
	GetAllUsers() (*[]models.UserDb, error)
	GetUserById(int) (*models.UserDb, error)
}
