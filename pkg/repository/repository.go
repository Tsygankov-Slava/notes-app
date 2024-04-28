package repository

import (
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GetUser(username, password string) (notes.User, error)
}

type Notes interface {
	Create(userId int, note notes.Note) (int, error)
	GetAll() ([]notes.Note, error)
	GetById(id int) (notes.Note, error)
	GetByUserId(userId int) ([]notes.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, input notes.UpdateNoteInput) error
}

type Repository struct {
	Authorization
	Notes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Notes:         NewNotePostgres(db),
	}
}
