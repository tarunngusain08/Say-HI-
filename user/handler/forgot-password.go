package handler

import (
	"Say-Hi/user/external"
	"Say-Hi/user/service"
	"Say-Hi/user/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ForgotPasswordHandler struct {
	forgotPasswordService *service.ForgotPasswordService
	emailService          *external.EmailService
}

func NewForgotPasswordHandler(service *service.ForgotPasswordService, emailService *external.EmailService) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		forgotPasswordService: service,
		emailService:          emailService,
	}
}

func (f *ForgotPasswordHandler) ForgotPassword(c *gin.Context) {

	user, err := validators.ValidateForgotPasswordUserDetails(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "StatusBadRequest")
	}

	err = f.forgotPasswordService.ForgotPassword(f.emailService, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "StatusInternalServerError")
	}

	c.JSON(http.StatusOK, "StatusOK")
}
