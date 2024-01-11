package handler

import (
	"Say-Hi/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type LogoutHandler struct {
	service *service.LogoutService
}

func NewLogoutHandler(service *service.LogoutService) *LogoutHandler {
	return &LogoutHandler{
		service: service,
	}
}

func (l *LogoutHandler) Logout(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, "Session Expired, logged out already")
		return
	}

	tokenParts := strings.Split(tokenString, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, "Invalid Authorization header format")
		return
	}

	jwtToken := tokenParts[1]
	err := l.service.Logout(jwtToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to logout, Something went wrong")
	}

	c.JSON(http.StatusOK, "Logged out")
}
