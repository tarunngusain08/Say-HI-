package repo

import (
	"Say-Hi/user/constants"
	"Say-Hi/user/contracts"
	"database/sql"
)

type RegisterRepo struct {
	db *sql.DB
}

func NewRegisterRepo(db *sql.DB) *RegisterRepo {
	return &RegisterRepo{db: db}
}

var register = `Insert into users (username, name, password, email, status, otp) values ($1,$2,$3,$4,$5,$6)`

func (r *RegisterRepo) Register(data *contracts.RegisterUser, otp string) error {
	_, err := r.db.Exec(register, data.UserName, data.Name, data.Password, data.Email, constants.Unverified, otp)
	if err != nil {
		return err
	}
	return nil
}
