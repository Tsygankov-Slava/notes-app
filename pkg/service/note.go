package service

import (
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/Tsygankov-Slava/notes-app/pkg/repository"
)

type NoteService struct {
	repo repository.Notes
}

func NewNoteService(repo repository.Notes) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(userId int, note notes.Note) (int, error) {
	id, err := s.repo.Create(userId, note)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *NoteService) GetAll() ([]notes.Note, error) {
	noteList, err := s.repo.GetAll()
	return noteList, err
}

func (s *NoteService) GetById(id int) (notes.Note, error) {
	note, err := s.repo.GetById(id)
	return note, err
}

//func (s *NoteService) GetByUserId(userId int) ([]notes.Note, error) {
//
//}

func (s *NoteService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *NoteService) Update(userId, noteId int, input notes.UpdateNoteInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, noteId, input)
}
