package validators

import (
	"Say-Hi/user/contracts"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"regexp"
)

func ValidateRegisterUserDetails(c *gin.Context) (*contracts.RegisterUser, error) {
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

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	ok := emailRegex.MatchString(user.Email)
	if !ok {
		return nil, errors.New("invalid email")
	}

	return &user, nil // Return the address of the user variable
}

func ValidateLoginUserDetails(c *gin.Context) (*contracts.LoginUser, error) {
	if c.Request == nil {
		return nil, errors.New("request error")
	}
	body := c.Request.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var user contracts.LoginUser
	err = json.Unmarshal(data, &user) // Pass the address of the user variable

	if err != nil {
		return nil, err
	}

	if (user.Email == "" && user.UserName == "") || user.Password == "" {
		return nil, errors.New("fill all the mandatory fields")
	}

	if user.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		ok := emailRegex.MatchString(user.Email)
		if !ok {
			return nil, errors.New("invalid email")
		}
	}

	return &user, nil // Return the address of the user variable
}
