package repository

import (
	"fmt"
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(user notes.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", usersTable)
	row := p.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *AuthPostgres) GetUser(username, password string) (notes.User, error) {
	var user notes.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := p.db.Get(&user, query, username, password)
	return user, err
}
