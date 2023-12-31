package middleware

import (
	"Say-Hi/user/contracts"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
)

func ValidateUserDetails(c *gin.Context) (*contracts.RegisterUser, error) {
	if c.Request == nil {
		return nil, errors.New("request error")
	}
	body := c.Request.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var user contracts.RegisterUser
	err = json.Unmarshal(data, &user) // Pass the address of the user variable

	if err != nil {
		return nil, err
	}

	if user.Name == "" || user.Email == "" || user.Password == "" || user.UserName == "" {
		return nil, errors.New("fill all the mandatory fields")
	}

	return &user, nil // Return the address of the user variable
}
