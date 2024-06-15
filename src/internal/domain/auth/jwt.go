package auth

// create token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"time"
)

func (a *AuthSrv) generateToken(user *models.UserDb, verified bool) (string, error) {
	claims := &models.Claims{
		UserId:    user.Id,
		Email:     user.Email,
		RoleId:    user.RoleId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Verified:  verified,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(utils.GetEnv("JWT_SECRET"))
}

func (a *AuthSrv) verifyToken(tokenString string) (*models.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.Log.Error("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return utils.GetEnv("JWT_SECRET"), nil
	})

	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	} else {
		utils.Log.Error("invalid token")
		return nil, fmt.Errorf("invalid token")
	}
}
