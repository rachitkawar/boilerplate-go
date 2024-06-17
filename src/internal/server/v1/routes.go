package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain"
)

type V1 struct {
	//all the domain mapping here
	srv          *domain.Service
	socialLogins *SocialLogin
}

func InitializeV1Routes(api *gin.RouterGroup, srv *domain.Service) {

	v1 := V1{
		srv:          srv,
		socialLogins: NewSocialLogins(srv.Auth),
	}
	v1.registerRoutesV1(api)
}

func (router *V1) registerRoutesV1(api *gin.RouterGroup) {
	rg := api.Group("/v1")

	auth := rg.Group("/auth")
	{
		auth.POST("/signup", router.Signup)
		auth.POST("/login", router.Login)
		auth.POST("/verify", router.Verify)
		auth.POST("/logout", router.Logout)

		socialLogins := auth.Group("/socialLogin")
		{
			google := socialLogins.Group("/google")
			{
				google.GET("/login", router.socialLogins.Google.SocialLogin)

				google.GET("/callback", router.socialLogins.Google.SocialLoginCallback)
			}
		}
	}

	roles := rg.Group("/roles")
	{
		roles.GET("/list")
		roles.POST("/create")
		roles.POST("/update")
		roles.POST("/delete")
	}

}
