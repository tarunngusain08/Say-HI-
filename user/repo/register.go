package repo

import (
	"Say-Hi/user/contracts"
	"database/sql"
)

type RegisterRepo struct {
	db *sql.DB
}

func NewRegisterRepo(db *sql.DB) *RegisterRepo {
	return &RegisterRepo{db: db}
}

var register = `Insert into Users values ($1,$2,$3,$4)`

func (r *RegisterRepo) Register(data *contracts.RegisterUser) error {
	_, err := r.db.Exec(register, data)
	if err != nil {
		return err
	}
	return nil
}
