package database

// permissions CRUD Operations

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
)

func (d *DB) GetAllPermissions() (*[]models.PermissionsDb, error) {
	query := `select * from permissions`
	result, err := d.db.Query(d.ctx, query)
	defer result.Close()
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query permissions: %w", err)
	}

	permissions, err := pgx.CollectRows(result, pgx.RowToStructByName[models.PermissionsDb])
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for permissions: %w", err)
	}
	return &permissions, err
}

func (d *DB) GetPermissionById(Id int) (*models.PermissionsDb, error) {
	query := `select * from permissions where id = @id`
	args := pgx.NamedArgs{
		"id": Id,
	}
	result, err := d.db.Query(d.ctx, query, args)
	defer result.Close()
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query permissions: %w", err)
	}

	permission, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[models.PermissionsDb])
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for permissions: %w", err)
	}
	return &permission, err
}

func (d *DB) GetPermissionByName(name string) (*models.PermissionsDb, error) {
	query := `select * from permissions where name = @name`
	args := pgx.NamedArgs{
		"name": name,
	}
	result, err := d.db.Query(d.ctx, query, args)
	defer result.Close()
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query permissions: %w", err)
	}

	permission, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[models.PermissionsDb])
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for permissions: %w", err)
	}
	return &permission, err
}

func (d *DB) DeletePermission(Id int) error {
	query := `delete from permissions where id = @id`
	args := pgx.NamedArgs{
		"id": Id,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to delete permission: %w", err)
	}
	return nil
}

func (d *DB) UpdatePermission(permission *models.PermissionsDb) error {
	query := `update permissions set name = @name , role_id = @role_id where id = @id`
	args := pgx.NamedArgs{
		"name":    permission.Name,
		"role_id": permission.RoleId,
		"id":      permission.Id,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to update permission: %w", err)
	}
	return nil

}

func (d *DB) CreatePermission(permission *models.PermissionsDb) error {
	query := `insert into permissions (name , role_id , created_at) 
				values (@name , @role_id , @created_at)`
	args := pgx.NamedArgs{
		"name":       permission.Name,
		"role_id":    permission.RoleId,
		"created_at": permission.CreatedAt,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to add permission: %w", err)
	}
	return nil
}

func (d *DB) CheckPermissionOnRoleId(roleId int, permissionName string) (bool, error) {
	query := `select exists(select 1 from permissions where role_id = @role_id and name = @name)`
	args := pgx.NamedArgs{
		"role_id": roleId,
		"name":    permissionName,
	}
	result, err := d.db.Query(d.ctx, query, args)
	defer result.Close()
	if err != nil {
		utils.Log.Error(err)
		return false, fmt.Errorf("unable to query permissions: %w", err)
	}

	var exists bool
	if result.Next() {
		err = result.Scan(&exists)
		if err != nil {
			utils.Log.Error(err)
			return false, fmt.Errorf("unable to scan the query for permissions: %w", err)
		}
	}
	return exists, nil
}
