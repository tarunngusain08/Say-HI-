package handler

import (
	"Say-Hi/user/middleware"
	"Say-Hi/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {

	err := middleware.ValidateUserDetails(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "StatusBadRequest")
	}

	err = service.Register()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "StatusInternalServerError")
	}

	c.JSON(http.StatusOK, "StatusOK")
}
