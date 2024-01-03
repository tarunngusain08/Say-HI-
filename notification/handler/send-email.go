package handler

import (
	"Say-Hi/notification/middleware"
	"Say-Hi/notification/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendEmailHandler struct {
	emailService *service.SendEmailService
}

func NewSendEmailHandler(emailService *service.SendEmailService) *SendEmailHandler {
	return &SendEmailHandler{
		emailService: emailService,
	}
}

func (s *SendEmailHandler) SendEmail(c *gin.Context) {

	fmt.Println("Calling SendEmail API")
	data, err := middleware.ValidateEmailDetails(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = s.emailService.SendEmail(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Success")
}
