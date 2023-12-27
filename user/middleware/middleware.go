package middleware

import (
	"Say-Hi/user"
	"Say-Hi/user/repo"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
)

func ValidateUserDetails(c *gin.Context) error {
	body, err := c.Request.GetBody()
	if err != nil {
		return err
	}
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	var User user.Customer
	err = json.Unmarshal(data, &User)

	if User.GetUserName() == "" || User.GetEmail() == "" || User.GetPassword() == "" || User.GetName() == "" {
		return errors.New("fill all the mandatory fields")
	}

	userNameExists := func() bool {
		return repo.UserNameExists()
	}()

	if userNameExists {
		return errors.New("username exists")
	}

	emailExists := func() bool {
		return repo.EmailExists()
	}()

	if emailExists {
		return errors.New("email exists, use different email or login")
	}
	return nil
}
