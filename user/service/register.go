package service

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/external"
	"Say-Hi/user/repo"
)

type RegisterService struct {
	repo *repo.RegisterRepo
}

func NewRegisterService(registerRepo *repo.RegisterRepo) *RegisterService {
	return &RegisterService{repo: registerRepo}
}

func (r *RegisterService) Register(service *external.EmailService, data *contracts.RegisterUser) error {

	otp, err := service.SendEmailWithExponentialBackoff(data.Email, data.Name)
	if err != nil {
		return err
	}

	err = r.repo.Register(data, *otp)
	if err != nil {
		return err
	}

	return nil
}
