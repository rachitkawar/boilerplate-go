package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain/auth"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"net/http"
)

type AuthMiddleware interface {
	AuthorizeToken() gin.HandlerFunc
}

// create a middleware to verify the token and it should be a function of router class

type UserAuthMiddleware struct {
	authSrv *auth.AuthSrv
}

func NewUserAuthMiddleware(authSrv *auth.AuthSrv) AuthMiddleware {
	return &UserAuthMiddleware{authSrv: authSrv}
}

func (u *UserAuthMiddleware) AuthorizeToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Unauthorized", nil, fmt.Errorf("no token provided")))
			c.Abort()
			return
		}

		token := &models.VerifyRequest{Token: tokenString}
		// Verify the token
		unwrappedToken, err := u.authSrv.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Invalid token", nil, err))
			c.Abort()
			return
		}

		// Set the user ID in the context
		c.Set("user_id", unwrappedToken.UserId)

		// Continue to the next handler
		c.Next()
	}
}
