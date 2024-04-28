package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type statusResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(context *gin.Context, status int, message string) {
	logrus.Error(message)
	context.AbortWithStatusJSON(status, errorResponse{Message: message}) // Blocking the execution of subsequent handlers
}
