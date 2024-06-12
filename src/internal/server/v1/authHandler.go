package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (router *V1) TestHandler(c *gin.Context) {

	router.srv.Auth.Login()
	c.JSON(http.StatusOK, gin.H{"message": "heelo"})
}
