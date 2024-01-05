package service

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/repo"
	"errors"
)

type LoginService struct {
	repo *repo.LoginRepo
}

func NewLoginRepo(loginRepo *repo.LoginRepo) *LoginService {
	return &LoginService{repo: loginRepo}
}

func (l *LoginService) Login(userDetails *contracts.LoginUser) error {
	if userDetails.UserName != "" {
		password, err := l.repo.GetUserPasswordByUsername(userDetails.UserName)
		if err != nil {
			return err
		}
		if password != userDetails.Password {
			return errors.New("wrong password")
		}
	}
	if userDetails.Email != "" {
		username, password, err := l.repo.GetUserPasswordByEmail(userDetails.Email)
		if err != nil {
			return err
		}
		if password != userDetails.Password {
			return errors.New("wrong password")
		}
		userDetails.UserName = username
	}
	return nil
}
