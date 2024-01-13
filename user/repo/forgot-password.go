package repo

import (
	"database/sql"
	"errors"
)

type ForgotPasswordRepo struct {
	db *sql.DB
}

func NewForgotPasswordRepo(db *sql.DB) *ForgotPasswordRepo {
	return &ForgotPasswordRepo{db: db}
}

var updateOTP = `UPDATE users SET otp = $1 WHERE email = $2;`

func (f *ForgotPasswordRepo) ForgotPassword(otp, email string) error {
	result, err := f.db.Exec(updateOTP, otp, email)
	if err != nil {
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("OTP was not updated in the DB")
	}

	return nil
}
