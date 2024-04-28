package repository

import (
	"fmt"
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/jmoiron/sqlx"
	"strings"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (p *NotePostgres) Create(userId int, note notes.Note) (int, error) {
	/* Create transaction */
	tx, err := p.db.Beginx()
	if err != nil {
		return 0, err
	}

	var id int
	// Query return notes id
	createNoteQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", notesTable)
	row := tx.QueryRow(createNoteQuery, note.Title, note.Description)
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback() // Rolling back the changes
		return 0, err
	}
	createUsersNotesQuery := fmt.Sprintf("INSERT INTO %s (user_id, notes_id) VALUES ($1, $2)", usersNotesTable)
	_, err = tx.Exec(createUsersNotesQuery, userId, id)
	if err != nil {
		err := tx.Rollback() // Rolling back the changes
		return 0, err
	}

	return id, tx.Commit()
}

func (p *NotePostgres) GetAll() ([]notes.Note, error) {
	var notesList []notes.Note
	query := fmt.Sprintf("SELECT * FROM %s", notesTable)
	rows, err := p.db.Query(query)
	if err != nil {
		return notesList, err
	}
	for rows.Next() {
		var note notes.Note
		if err := rows.Scan(&note.Id, &note.Title, &note.Description); err != nil {
			return notesList, err
		}
		notesList = append(notesList, note)
	}
	if err = rows.Err(); err != nil {
		return notesList, err
	}
	return notesList, nil
}

func (p *NotePostgres) GetById(id int) (notes.Note, error) {
	var note notes.Note
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", notesTable)
	err := p.db.Get(&note, query, id)
	return note, err
}

func (p *NotePostgres) GetByUserId(userId int) ([]notes.Note, error) {
	var notesList []notes.Note
	query := fmt.Sprintf("SELECT nt.id, nt.title, nt.description FROM %s nt INNER JOIN %s unt on nt.id = unt.notes_id WHERE unt.user_id = $1",
		notesTable, usersNotesTable)
	err := p.db.Select(&notesList, query, userId)
	return notesList, err
}

func (p *NotePostgres) Delete(userId, noteId int) error {
	query := fmt.Sprintf("DELETE FROM %s nt USING %s unt WHERE nt.id = unt.notes_id AND unt.user_id = $1 AND unt.notes_id=$2",
		notesTable, usersNotesTable)
	_, err := p.db.Exec(query, userId, noteId)
	return err
}

func (p *NotePostgres) Update(userId, noteId int, input notes.UpdateNoteInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s nt SET %s FROM %s unt WHERE nt.id = unt.notes_id AND unt.notes_id=$%d AND unt.user_id = $%d`,
		notesTable, setQuery, usersNotesTable, argId, argId+1)

	args = append(args, noteId, userId)

	_, err := p.db.Exec(query, args...)
	return err
}
