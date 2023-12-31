package handler

import (
	"Say-Hi/user/middleware"
	"Say-Hi/user/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterHandler struct {
	repo *repo.RegisterRepo
}

func NewRegisterHandler(registerRepo *repo.RegisterRepo) *RegisterHandler {
	return &RegisterHandler{repo: registerRepo}
}

func (r *RegisterHandler) Register(c *gin.Context) {

	data, err := middleware.ValidateUserDetails(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = r.repo.Register(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Success")
}
