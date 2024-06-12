package database

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
)

func (d *DB) GetAllUsers() (*[]models.UserDb, error) {
	query := `select * from users`
	result, err := d.db.Query(d.ctx, query)
	defer result.Close()
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	users, err := pgx.CollectRows(result, pgx.RowToStructByName[models.UserDb])
	if err != nil {
		utils.Log.Error(err)
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
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	user, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[models.UserDb])
	if err != nil {
		utils.Log.Error(err)
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
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	user, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[models.UserDb])
	if err != nil {
		utils.Log.Error(err)
		return nil, fmt.Errorf("unable to scan the query for users: %w", err)
	}
	return &user, err
}

func (d *DB) AddUser(user *models.UserDb) error {
	query := `insert into users (first_name , last_name, email, phone_number ,password ,role_id ,  created_at) 
				values (@first_name , @last_name, @email, @phone_number ,@password ,@role_id, @created_at)`
	args := pgx.NamedArgs{
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"password":     user.Password,
		"role_id":      user.RoleId,
		"created_at":   user.CreatedAt,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to add user: %w", err)
	}
	return nil
}

// UpdateUser cannot update User Role
func (d *DB) UpdateUser(user *models.UserDb) error {
	query := `update users set first_name = @first_name , last_name = @last_name, email = @email, 
                 phone_number = @phone_number, password = @users.password where id = @id`
	args := pgx.NamedArgs{
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"password":     user.Password,
		"id":           user.Id,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to update user: %w", err)
	}
	return nil

}

func (d *DB) DeleteUser(Id int) error {
	query := `delete from users where id = @id`
	args := pgx.NamedArgs{
		"id": Id,
	}
	_, err := d.db.Exec(d.ctx, query, args)
	if err != nil {
		utils.Log.Error(err)
		return fmt.Errorf("unable to delete user: %w", err)
	}
	return nil
}
