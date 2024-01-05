package repo

import (
	"database/sql"
	"errors"
)

type LoginRepo struct {
	db *sql.DB
}

func NewLoginRepo(db *sql.DB) *LoginRepo {
	return &LoginRepo{db: db}
}

var getUserDetailsByUsername = `Select password from users where userName = $1`

var getUserDetailsByEmail = `Select username, password from users where email = $1`

func (repo *LoginRepo) GetUserPasswordByUsername(username string) (string, error) {
	var password string
	err := repo.db.QueryRow(getUserDetailsByUsername, username).Scan(&password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("userName not found")
		}
		return "", err
	}
	return password, nil
}

func (repo *LoginRepo) GetUserPasswordByEmail(email string) (string, string, error) {
	var username, password string
	err := repo.db.QueryRow(getUserDetailsByEmail, email).Scan(&username, &password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", errors.New("email not found")
		}
		return "", "", err
	}
	return username, password, nil
}
