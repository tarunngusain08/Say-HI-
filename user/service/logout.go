package service

import (
	"Say-Hi/auth"
	"Say-Hi/user/repo"
)

type LogoutService struct {
	repo *repo.LogoutRepo
}

func NewLogoutService(logoutRepo *repo.LogoutRepo) *LogoutService {
	return &LogoutService{repo: logoutRepo}
}

func (l *LogoutService) Logout(jwtToken string) error {

	auth.AddToBlacklist(jwtToken)
	return nil
}
