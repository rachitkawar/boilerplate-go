package v1

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain/auth"
	"github.com/rachitkawar/boilerplate-go/src/internal/models"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
	"time"
)

type SocialLogin struct {
	Google   OauthLogins
	Github   OauthLogins
	LinkedIn OauthLogins
}

func NewSocialLogins(auth *auth.AuthSrv) *SocialLogin {
	return &SocialLogin{
		Google: NewGoogleConfig(auth),
	}
}

type OauthLogins interface {
	SocialLogin(c *gin.Context)
	SocialLoginCallback(c *gin.Context)
}

type GoogleConfig struct {
	config oauth2.Config
	auth   *auth.AuthSrv
}

func NewGoogleConfig(auth *auth.AuthSrv) OauthLogins {
	googleLoginConfig := &GoogleConfig{
		config: oauth2.Config{
			RedirectURL:  utils.GetEnv("GOOGLE_REDIRECT_URL"),
			ClientID:     utils.GetEnv("GOOGLE_CLIENT_ID"),
			ClientSecret: utils.GetEnv("GOOGLE_CLIENT_SECRET"),
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint: google.Endpoint,
		},
		auth: auth,
	}

	return googleLoginConfig
}

func (gc *GoogleConfig) SocialLogin(c *gin.Context) {

	url := gc.config.AuthCodeURL(utils.GetEnv("SOCIAL_LOGIN_SECRET_STATE"))

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (gc *GoogleConfig) SocialLoginCallback(c *gin.Context) {
	state := c.Query("state")
	if state != utils.GetEnv("SOCIAL_LOGIN_SECRET_STATE") {
		c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("GOOGLE_LOGIN_URL"))
		return
	}

	code := c.Query("code")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	token, err := gc.config.Exchange(ctx, code)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("GOOGLE_LOGIN_URL"))
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("GOOGLE_LOGIN_URL"))
		return
	}

	// get userGoogleLoginModel from resp
	defer resp.Body.Close()
	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("GOOGLE_LOGIN_URL"))
		return
	}

	userGoogleLoginRequest := &models.GoogleLoginRequest{}
	err = json.Unmarshal(userData, userGoogleLoginRequest)
	if err != nil {
		utils.Log.Error("cannot load this model with json (", string(userData), " )")
		c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("GOOGLE_LOGIN_URL"))
		return
	}

	// generate token
	socialLoginModel := &models.SocialLoginModel{
		Email:     userGoogleLoginRequest.Email,
		FirstName: userGoogleLoginRequest.GivenName,
		LastName:  userGoogleLoginRequest.FamilyName,
	}

	loginToken, err := gc.auth.SocialLogin(socialLoginModel)
	if err != nil {
		utils.Log.Error("cannot load this model with json (", string(userData), " )")
		c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("GOOGLE_LOGIN_URL"))
		return
	}

	c.SetCookie("authToken", loginToken.Token, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusPermanentRedirect, "http://google.com")

}
