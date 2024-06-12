package auth

import (
	"fmt"
	"github.com/rachitkawar/boilerplate-go/src/internal/database"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
)

type AuthSrv struct {
	db database.Store
}

func NewAuthSrv(db database.Store) *AuthSrv {
	return &AuthSrv{
		db: db,
	}
}

func (a *AuthSrv) Login() (*models.LoginResponse, error) {
	fmt.Println("here login")

	users, err := a.db.GetAllUsers()
	fmt.Println(users, err)

	// if login success

	//token, err := a.generateToken(user)
	//if err != nil {
	//	utils.Log.Error("cannot generate token %v", err)
	//	return &models.LoginResponse{}, fmt.Errorf("cannot generate login token")
	//}
	// if login failed

	return &models.LoginResponse{}, nil
}
