package handler

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type VerifyEmailHandler struct {
	verifyEmailService *service.VerifyEmailService
}

func NewVerifyEmailHandler(emailService *service.VerifyEmailService) *VerifyEmailHandler {
	return &VerifyEmailHandler{
		verifyEmailService: emailService,
	}
}

func (v *VerifyEmailHandler) VerifyEmail(c *gin.Context) {

	fmt.Println("Calling VerifyEmail API")
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var body contracts.VerifyEmailRequest
	err = json.Unmarshal(data, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = v.verifyEmailService.VerifyEmail(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "email verified")
}
