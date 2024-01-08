package handler

import (
	"Say-Hi/auth"
	"Say-Hi/user/service"
	"Say-Hi/user/validators"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	service *service.LoginService
	jwt     *auth.JWT
}

func NewLoginHandler(service *service.LoginService, jwt *auth.JWT) *LoginHandler {
	return &LoginHandler{
		service: service,
		jwt:     jwt,
	}
}

func (l *LoginHandler) Login(c *gin.Context) {

	fmt.Println("Calling Login API")
	userDetails, err := validators.ValidateLoginUserDetails(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid login details")
		return
	}

	err = l.service.Login(userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to login, Something went wrong")
		return
	}

	token, err := l.jwt.GenerateJWT(userDetails.UserName, userDetails.Password)

	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, "Logged in")
}
