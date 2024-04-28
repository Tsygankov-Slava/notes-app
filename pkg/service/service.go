package service

import (
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/Tsygankov-Slava/notes-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Notes interface {
	Create(userId int, note notes.Note) (int, error)
	GetAll() ([]notes.Note, error)
	GetById(id int) (notes.Note, error)
	//GetByUserId(userId int) ([]notes.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, input notes.UpdateNoteInput) error
}

type Service struct {
	Authorization
	Notes
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Notes:         NewNoteService(repo.Notes),
	}
}
