package handler

import (
	"Say-Hi/auth"
	"Say-Hi/user/middleware"
	"Say-Hi/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	service *service.LoginService
	jwt     *auth.JWT
}

func NewLoginRepo(service *service.LoginService, jwt *auth.JWT) *LoginHandler {
	return &LoginHandler{
		service: service,
		jwt:     jwt,
	}
}

func (l *LoginHandler) Login(c *gin.Context) {

	userDetails, err := middleware.ValidateLoginUserDetails(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "StatusBadRequest")
	}

	err = l.service.Login(userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "StatusInternalServerError")
	}

	token, err := l.jwt.GenerateJWT(userDetails.UserName, userDetails.Password)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
