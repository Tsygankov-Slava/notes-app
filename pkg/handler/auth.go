package handler

import (
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body notes.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(context *gin.Context) {
	var input notes.User
	if err := context.BindJSON(&input); err != nil {
		/* Incorrect data is specified in the request */
		newErrorResponse(context, http.StatusBadRequest /* 400 */, err.Error())
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		/* Internal server error */
		newErrorResponse(context, http.StatusInternalServerError /* 500 */, err.Error())
		return
	}

	context.JSON(http.StatusOK /* 200 */, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(context *gin.Context) {
	var input signInInput
	if err := context.BindJSON(&input); err != nil {
		/* Incorrect data is specified in the request */
		newErrorResponse(context, http.StatusBadRequest /* 400 */, err.Error())
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		/* Internal server error */
		newErrorResponse(context, http.StatusInternalServerError /* 500 */, err.Error())
		return
	}

	context.JSON(http.StatusOK /* 200 */, map[string]interface{}{
		"token": token,
	})
}
