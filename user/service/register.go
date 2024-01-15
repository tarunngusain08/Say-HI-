package service

import (
	"Say-Hi/user/contracts"
	"Say-Hi/user/external"
	"Say-Hi/user/repo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	repo *repo.RegisterRepo
}

func NewRegisterService(registerRepo *repo.RegisterRepo) *RegisterService {
	return &RegisterService{repo: registerRepo}
}

func (r *RegisterService) Register(service *external.EmailService, data *contracts.UserDetails) error {

	hashedPassword, err := r.hashPassword(data.Password)
	if err != nil {
		return err
	}
	data.Password = hashedPassword

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

// hashPassword hashes the provided password using bcrypt
func (r *RegisterService) hashPassword(password string) (string, error) {
	// Generate a salted hash for the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
