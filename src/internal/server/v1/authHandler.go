package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"net/http"
)

// Signup godoc
//
//	@Summary		User Signup
//	@Tags			Signup
//	@Accept			json
//	@Produce		json
//	@Param			SignupRequest	body		models.SignupRequest	true	"Add user"
//	@Success		200		{object}	utils.ApiResponse
//	@Failure		400		{object}    utils.ApiResponse
//	@Failure		404		{object}	utils.ApiResponse
//	@Failure		500		{object}	utils.ApiResponse
//	@Router			/api/v1/auth/signup [post]
func (router *V1) Signup(c *gin.Context) {
	var signupRequest models.SignupRequest
	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Invalid request", nil, err))
		return
	}

	signupResponse, err := router.srv.Auth.Signup(&signupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Internal server error", nil, err))
		return
	}

	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Success", signupResponse, nil))

}

// Login godoc
//
//	@Summary		User Login
//	@Tags			Login
//	@Accept			json
//	@Produce		json
//	@Param			LoginRequest	body		models.LoginRequest	true	"Login user"
//	@Success		200		{object}	utils.ApiResponse
//	@Failure		400		{object}    utils.ApiResponse
//	@Failure		404		{object}	utils.ApiResponse
//	@Failure		500		{object}	utils.ApiResponse
//	@Router			/api/v1/auth/login [post]

func (router *V1) Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Invalid request", nil, err))
		return
	}

	loginResponse, err := router.srv.Auth.Login(&loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Internal server error", nil, err))
		return
	}

	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Success", loginResponse, nil))
}

//
// Verify godoc
//
//	@Summary		Verify Token
//	@Tags			Verify
//	@Accept			json
//	@Produce		json
//	@Param			VerifyRequest	body		models.VerifyRequest	true	"Verify token"
//	@Success		200		{object}	utils.ApiResponse
//	@Failure		400		{object}    utils.ApiResponse
//	@Failure		404		{object}	utils.ApiResponse
//	@Failure		500		{object}	utils.ApiResponse
//	@Router			/api/v1/auth/verify [post]

func (router *V1) Verify(c *gin.Context) {
	var verifyRequest models.VerifyRequest
	if err := c.ShouldBindJSON(&verifyRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Invalid request", nil, err))
		return
	}

	verifyResponse, err := router.srv.Auth.VerifyToken(&verifyRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewApiResponse(utils.ApiResponseCode.ERROR, "Internal server error", nil, err))
		return
	}

	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Success", verifyResponse, nil))

}

func (router *V1) Logout(c *gin.Context) {
	//TODO: implement logout logic
	c.JSON(http.StatusOK, utils.NewApiResponse(utils.ApiResponseCode.SUCCESS, "Logout successful", nil, nil))

}
