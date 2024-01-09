package handler

import (
	"Say-Hi/user/external"
	"Say-Hi/user/service"
	"Say-Hi/user/validators"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterHandler struct {
	registerService *service.RegisterService
	emailService    *external.EmailService
}

func NewRegisterHandler(registerService *service.RegisterService, emailService *external.EmailService) *RegisterHandler {
	return &RegisterHandler{
		registerService: registerService,
		emailService:    emailService,
	}
}

func (r *RegisterHandler) Register(c *gin.Context) {

	fmt.Println("Calling Register API")
	data, err := validators.ValidateRegisterUserDetails(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = r.registerService.Register(r.emailService, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Success")
}
