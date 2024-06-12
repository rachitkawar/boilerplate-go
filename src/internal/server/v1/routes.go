package v1

import (
	"github.com/gin-gonic/gin"
)

type V1 struct {
	//all the domain mapping here
}

func InitializeV1Routes(api *gin.RouterGroup) {

	v1 := V1{}
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
