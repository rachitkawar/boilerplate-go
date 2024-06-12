package database

// Roles Crud Operations

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
)

func (d *DB) GetAllRoles() (*[]models.RolesDb, error) {
	query := `select * from roles`
	result, err := d.db.Query(d.ctx, query)
	defer result.Close()
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query roles: %w", err)
	}

	roles, err := pgx.CollectRows(result, pgx.RowToStructByName[models.RolesDb])
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for roles: %w", err)
	}
	return &roles, err
}

func (d *DB) GetRoleById(Id int) (*models.RolesDb, error) {
	query := `select * from roles where id = @id`
	args := pgx.NamedArgs{
		"id": Id,
	}
	result, err := d.db.Query(d.ctx, query, args)
	defer result.Close()
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query roles: %w", err)
	}

	role, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[models.RolesDb])
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for roles: %w", err)
	}
	return &role, err
}

func (d *DB) AddRole(role *models.RolesDb) error {
	query := `insert into roles (name , created_at) 
				values (@name , @created_at)`
	args := pgx.NamedArgs{
		"name":       role.Name,
		"created_at": role.CreatedAt,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to add role: %w", err)
	}
	return nil
}

func (d *DB) DeleteRole(Id int) error {
	query := `delete from roles where id = @id`
	args := pgx.NamedArgs{
		"id": Id,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to delete role: %w", err)
	}
	return nil
}

func (d *DB) UpdateRole(role *models.RolesDb) error {
	query := `update roles set name = @name where id = @id`
	args := pgx.NamedArgs{
		"name": role.Name,
		"id":   role.Id,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to update role: %w", err)
	}
	return nil

}
