package middleware

import (
	"Say-Hi/notification/contracts"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
)

func ValidateEmailDetails(c *gin.Context) (*contracts.SendEmailRequest, error) {
	if c.Request == nil {
		return nil, errors.New("request is empty")
	}
	body := c.Request.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var emailRequest contracts.SendEmailRequest
	err = json.Unmarshal(data, &emailRequest) // Pass the address of the emailRequest variable

	if err != nil {
		return nil, err
	}
	return &emailRequest, nil
}
