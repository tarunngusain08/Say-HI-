package middleware

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/repo"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
)

func ValidateUserDetails(c *gin.Context) (*contracts.RegisterUser, error) {
	body, err := c.Request.GetBody()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var user *contracts.RegisterUser
	err = json.Unmarshal(data, user)

	if user.Name == "" || user.Email == "" || user.Password == "" || user.Name == "" {
		return nil, errors.New("fill all the mandatory fields")
	}

	userNameExists := func() bool {
		return repo.UserNameExists()
	}()

	if userNameExists {
		return nil, errors.New("username exists")
	}

	emailExists := func() bool {
		return repo.EmailExists()
	}()

	if emailExists {
		return nil, errors.New("email exists, use different email or login")
	}
	return user, nil
}
