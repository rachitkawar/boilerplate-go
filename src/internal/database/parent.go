package database

import "github.com/rachitkawar/boilerplate-go/src/internal/models"

type Store interface {
	Close()

	GetAllUsers() (*[]models.UserDb, error)
	GetUserById(int) (*models.UserDb, error)
	CreateUser(user *models.UserDb) error
	UpdateUser(user *models.UserDb) error
	DeleteUser(int) error

	GetAllRoles() (*[]models.RolesDb, error)
	GetRoleById(int) (*models.RolesDb, error)
	CreateRole(*models.RolesDb) error
	UpdateRole(*models.RolesDb) error
	DeleteRole(int) error

	GetAllPermissions() (*[]models.PermissionsDb, error)
	GetPermissionById(int) (*models.PermissionsDb, error)
	CreatePermission(*models.PermissionsDb) error
	UpdatePermission(*models.PermissionsDb) error
	DeletePermission(int) error
	CheckPermissionOnRoleId(int, string) (bool, error)
}
