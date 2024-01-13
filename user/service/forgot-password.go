package service

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/external"
	"Say-Hi/user/repo"
)

type ForgotPasswordService struct {
	repo *repo.ForgotPasswordRepo
}

func NewForgotPasswordService(forgotPasswordRepo *repo.ForgotPasswordRepo) *ForgotPasswordService {
	return &ForgotPasswordService{repo: forgotPasswordRepo}
}

func (f *ForgotPasswordService) ForgotPassword(service *external.EmailService, user *contracts.UserDetails) error {
	otp, err := service.SendEmailWithExponentialBackoff(user.Email, user.Name)
	if err != nil {
		return err
	}

	err = f.repo.ForgotPassword(*otp, user.Email)
	if err != nil {
		return err
	}

	return nil
}
