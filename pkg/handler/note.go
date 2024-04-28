package handler

import (
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
