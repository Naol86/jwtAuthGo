package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naol86/go/authGo/config"
	"github.com/naol86/go/authGo/internal/domain"
	"github.com/naol86/go/authGo/package/tokens"
)

type UserController struct {
	UserUseCase domain.UserUseCase
	Env         *config.Env
}

func (uc *UserController) Signin(c *gin.Context) {
	var user domain.UserSigninRequest
	var response domain.UserResponse
	if err := c.ShouldBind(&user); err != nil {
		response.Message = "Invalid request"
		response.Success = false
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userData, err := uc.UserUseCase.Signin(c.Request.Context(), user)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	accessToken, err := tokens.CreateAccessToken(&userData, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpireHour)
	if err != nil {
		response.Message = "Failed to create access token"
		response.Success = false
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	refreshToken, err := tokens.CreateRefreshToken(&userData, uc.Env.RefreshTokenSecret, uc.Env.RefreshTokenExpireHour)
	if err != nil {
		response.Message = "Failed to create refresh token"
		response.Success = false
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "user signed in successfully"
	response.Success = true
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	response.Data = userData
	c.JSON(http.StatusOK, response)

}
