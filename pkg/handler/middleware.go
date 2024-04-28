package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) UserIdentity(context *gin.Context) {
	header := context.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(context, http.StatusUnauthorized /* 401 */, "empty Authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 { // Checking for the correct header
		newErrorResponse(context, http.StatusUnauthorized /* 401 */, "invalid authorization header")
		return
	}
	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(context, http.StatusUnauthorized /* 401 */, err.Error())
		return
	}
	/* Write it to the context so that in subsequent handlers we can have access to the userId */
	context.Set("userId", userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get("userId")
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError /* 500 */, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError /* 500 */, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
