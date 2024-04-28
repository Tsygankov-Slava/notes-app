package notes

import "errors"

type Note struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersNotes struct {
	Id     int
	UserId int
	NoteId int
}

type UpdateNoteInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description""`
}

func (i UpdateNoteInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
