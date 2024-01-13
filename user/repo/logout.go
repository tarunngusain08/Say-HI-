package repo

import "database/sql"

type LogoutRepo struct {
	db *sql.DB
}

func NewLogoutRepo(db *sql.DB) *LogoutRepo {
	return &LogoutRepo{db: db}
}

func (l *LogoutRepo) Logout() error {
	return nil
}
