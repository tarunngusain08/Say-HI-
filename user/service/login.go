package service

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/repo"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo *repo.LoginRepo
}

func NewLoginService(loginRepo *repo.LoginRepo) *LoginService {
	return &LoginService{repo: loginRepo}
}

func (l *LoginService) Login(userDetails *contracts.UserDetails) error {
	if userDetails.UserName != "" {
		password, err := l.repo.GetUserPasswordByUsername(userDetails.UserName)
		if err != nil {
			return err
		}
		err = l.checkPasswordHash(password, userDetails.Password)
		if err != nil {
			return err
		}
	} else if userDetails.Email != "" {
		username, password, err := l.repo.GetUserPasswordByEmail(userDetails.Email)
		if err != nil {
			return err
		}
		err = l.checkPasswordHash(password, userDetails.Password)
		if err != nil {
			return err
		}
		userDetails.UserName = username
	}
	return nil
}

// checkPasswordHash checks if the provided password matches the hashed password
func (l *LoginService) checkPasswordHash(password, hash string) error {
	// Compare the password with the hashed version
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
