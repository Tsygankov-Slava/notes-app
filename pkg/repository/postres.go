package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	notesTable      = "notes"
	usersNotesTable = "users_notes"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode))

	if err != nil {
		return nil, err
	}

	return db, nil
}
