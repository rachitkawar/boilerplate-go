package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"net/http"
)

func (router *V1) ListRoles(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Success", nil, nil))
}

func (router *V1) CreateRole(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Success", nil, nil))
}

func (router *V1) UpdateRole(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Success", nil, nil))
}

func (router *V1) DeleteRole(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Success", nil, nil))
}
