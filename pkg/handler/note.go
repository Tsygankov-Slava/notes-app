package handler

import (
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create note
// @Security ApiKeyAuth
// @Tags notes
// @Description create note
// @ID create-note
// @Accept  json
// @Produce  json
// @Param input body notes.Note true "note info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes [post]
func (h *Handler) createNote(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}
	var input notes.Note
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest /* 400 */, err.Error())
		return
	}
	id, err := h.service.Notes.Create(userId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError /* 500 */, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllNotesResponse struct {
	Data []notes.Note `json:"data"`
}

// @Summary Get All Notes
// @Security ApiKeyAuth
// @Tags notes
// @Description get all notes
// @ID get-all-notes
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllNotesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes [get]
func (h *Handler) getAllNotes(ctx *gin.Context) {
	notesList, err := h.service.Notes.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError /* 500 */, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, getAllNotesResponse{
		Data: notesList,
	})
}

// @Summary Get Note By Id
// @Security ApiKeyAuth
// @Tags notes
// @Description get note by id
// @ID get-note-by-id
// @Accept json
// @Produce json
// @Param id path int true "Note id"
// @Success 200 {object} notes.Note
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes/{id} [get]
func (h *Handler) getNoteById(ctx *gin.Context) {
	noteId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest /* 400 */, "invalid  id param")
		return
	}
	note, err := h.service.Notes.GetById(noteId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError /* 500 */, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, note)
}

func (h *Handler) getNotesByUserId(ctx *gin.Context) {

}

// @Summary Update Note By Id
// @Security ApiKeyAuth
// @Tags notes
// @Description update note by id
// @ID update-note-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Note ID"
// @Param input body notes.UpdateNoteInput true "note info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes/{id} [put]
func (h *Handler) updateNoteById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	noteId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest /* 400 */, "invalid  id param")
		return
	}

	var input notes.UpdateNoteInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest /* 400 */, err.Error())
	}

	err = h.service.Update(userId, noteId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError /* 500 */, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Note By Id
// @Security ApiKeyAuth
// @Tags notes
// @Description delete note by id
// @ID delete-note-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Note ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes/{id} [delete]
func (h *Handler) deleteNoteById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	noteId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest /* 400 */, "invalid  id param")
		return
	}

	err = h.service.Notes.Delete(userId, noteId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError /* 500 */, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
