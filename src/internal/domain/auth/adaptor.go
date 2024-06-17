package auth

import (
	"fmt"
	"github.com/rachitkawar/boilerplate-go/src/internal/database"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthSrv struct {
	db database.Store
}

func NewAuthSrv(db database.Store) *AuthSrv {
	return &AuthSrv{
		db: db,
	}
}

func (a *AuthSrv) Login(loginRequest *models.LoginRequest) (*models.LoginResponse, error) {
	utils.Log.Info("user logging in")

	//check if user exists
	user, err := a.db.GetUserByEmail(loginRequest.Email)
	if err != nil {
		utils.Log.Error("cannot find user %v", err)
		return &models.LoginResponse{}, fmt.Errorf("cannot find user")
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		utils.Log.Error("cannot compare password %v", err)
		return &models.LoginResponse{}, fmt.Errorf("cannot compare password")
	}

	//generate token
	token, err := a.generateToken(user, true, true)
	if err != nil {
		utils.Log.Error("cannot generate token %v", err)
		return &models.LoginResponse{}, fmt.Errorf("cannot generate login token")
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}

func (a *AuthSrv) VerifyToken(verifyRequest *models.VerifyRequest) (*models.VerifyResponse, error) {
	utils.Log.Info("verifying token")
	//verify token
	claims, err := a.verifyToken(verifyRequest.Token)
	if err != nil {
		utils.Log.Error("cannot verify token %v", err)
		return &models.VerifyResponse{}, fmt.Errorf("cannot verify token")
	}

	return &models.VerifyResponse{
		UserId:    claims.UserId,
		RoleId:    claims.RoleId,
		Email:     claims.Email,
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		Verified:  claims.Verified,
	}, nil

}

func (a *AuthSrv) Logout() (*models.LoginResponse, error) {
	utils.Log.Info("user logging out")
	//TODO: implement logout logic

	return &models.LoginResponse{}, nil
}

func (a *AuthSrv) Signup(signupRequest *models.SignupRequest) (*models.SignupResponse, error) {
	utils.Log.Info("user signing up")

	//check if user not exists
	exists, err := a.db.CheckUserByEmail(signupRequest.Email)
	if err != nil {
		utils.Log.Error("cannot check if user exists %v", err)
		return &models.SignupResponse{}, fmt.Errorf("cannot check if user exists")
	}

	if exists {
		utils.Log.Error("user already exists")
		return &models.SignupResponse{}, fmt.Errorf("user already exists")
	}

	// hash password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Log.Error("cannot hash password %v", err)
		return &models.SignupResponse{}, fmt.Errorf("cannot hash password")
	}

	user := &models.UserDb{
		Email:       signupRequest.Email,
		Password:    string(hashedPassword),
		FirstName:   signupRequest.FirstName,
		LastName:    signupRequest.LastName,
		PhoneNumber: signupRequest.PhoneNumber,
		RoleId:      utils.DefaultRoleId,
	}

	//save to db
	err = a.db.CreateUser(user)
	if err != nil {
		utils.Log.Error("cannot create user %v", err)
		return &models.SignupResponse{}, fmt.Errorf("cannot create user")
	}

	//generate token
	token, err := a.generateToken(user, false, true)
	if err != nil {
		utils.Log.Error("cannot generate token %v", err)
		return &models.SignupResponse{}, fmt.Errorf("cannot generate signup token")
	}

	return &models.SignupResponse{
		Token: token,
	}, nil
}

func (a *AuthSrv) SocialLogin(socialLoginRequest *models.SocialLoginModel) (*models.SignupResponse, error) {
	//check if user not exists
	exists, err := a.db.CheckUserByEmail(socialLoginRequest.Email)
	if err != nil {
		utils.Log.Error("cannot check if user exists %v", err)
		return &models.SignupResponse{}, fmt.Errorf("cannot check if user exists")
	}
	user := &models.UserDb{}
	var profileComplete bool
	if !exists {
		user = &models.UserDb{
			Email:     socialLoginRequest.Email,
			FirstName: socialLoginRequest.FirstName,
			LastName:  socialLoginRequest.LastName,
			RoleId:    utils.DefaultRoleId,
		}

		//save to db
		err = a.db.CreateUser(user)
		if err != nil {
			utils.Log.Error("cannot create user %v", err)
			return &models.SignupResponse{}, fmt.Errorf("cannot create user")
		}
		profileComplete = false
	} else {
		//check if user exists
		user, err = a.db.GetUserByEmail(socialLoginRequest.Email)
		if err != nil {
			utils.Log.Error("cannot find user %v", err)
			return &models.SignupResponse{}, fmt.Errorf("cannot find user")
		}

		if user.Password == "" {
			profileComplete = false
		}

	}

	// generate token
	token, err := a.generateToken(user, false, profileComplete)
	if err != nil {
		utils.Log.Error("cannot generate token %v", err)
		return &models.SignupResponse{}, fmt.Errorf("cannot generate signup token")
	}

	return &models.SignupResponse{
		Token: token,
	}, nil

}
