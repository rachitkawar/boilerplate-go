package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain/jwt"
)

type V1 struct {
	//all the domain mapping here
	jwt *jwt.TokenMaster
}

func InitializeV1Routes(api *gin.RouterGroup, jwt *jwt.TokenMaster) {

	v1 := V1{
		jwt: jwt,
	}
	v1.registerRoutesV1(api)
}

func (router *V1) registerRoutesV1(api *gin.RouterGroup) {
	rg := api.Group("/v1")

	auth := rg.Group("/auth")
	{
		auth.GET("/test", router.TestHandler)
		auth.POST("/signup")
		auth.POST("/login")
		auth.POST("/verify")
	}

}
