package service

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/repo"
)

type VerifyEmailService struct {
	repo *repo.VerifyEmailRepo
}

func NewVerifyEmailService(repo *repo.VerifyEmailRepo) *VerifyEmailService {
	return &VerifyEmailService{
		repo: repo,
	}
}

func (v *VerifyEmailService) VerifyEmail(data contracts.VerifyEmailRequest) error {
	return v.repo.VerifyEmail(data)
}
