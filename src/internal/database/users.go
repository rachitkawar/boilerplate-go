package database

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rachitkawar/boilerplate-go/src/common"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
)

func (d *DB) GetAllUsers() (*[]models.UserDb, error) {
	query := `select * from users`
	result, err := d.db.Query(d.ctx, query)
	defer result.Close()
	if err != nil {
		common.Log.Error(err)
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	users, err := pgx.CollectRows(result, pgx.RowToStructByName[models.UserDb])
	if err != nil {
		common.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for users: %w", err)
	}
	return &users, err
}

func (d *DB) GetUserById(Id int) (*models.UserDb, error) {
	query := `select * from users where id = @id`
	args := pgx.NamedArgs{
		"id": Id,
	}
	result, err := d.db.Query(d.ctx, query, args)
	defer result.Close()
	if err != nil {
		common.Log.Error(err)
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	user, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[models.UserDb])
	if err != nil {
		common.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for users: %w", err)
	}
	return &user, err
}

func (d *DB) GetUserByEmail(Email string) (*models.UserDb, error) {
	query := `select * from users where email = @email`
	args := pgx.NamedArgs{
		"email": Email,
	}
	result, err := d.db.Query(d.ctx, query, args)
	defer result.Close()
	if err != nil {
		common.Log.Error(err)
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	user, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[models.UserDb])
	if err != nil {
		common.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for users: %w", err)
	}
	return &user, err
}
