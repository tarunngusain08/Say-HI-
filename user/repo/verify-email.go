package repo

import (
	"Say-Hi/user/contracts"
	"database/sql"
	"errors"
)

type VerifyEmailRepo struct {
	db *sql.DB
}

func NewVerifyEmailRepo(db *sql.DB) *VerifyEmailRepo {
	return &VerifyEmailRepo{
		db: db,
	}
}

var getOTP = `UPDATE users SET status = 'verified' WHERE email = $1 AND otp = $2;`

func (r *VerifyEmailRepo) VerifyEmail(data contracts.VerifyEmailRequest) error {
	result, err := r.db.Exec(getOTP, data.Email, data.OTP)
	if err != nil {
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("wrong OTP given")
	}
	return nil
}
