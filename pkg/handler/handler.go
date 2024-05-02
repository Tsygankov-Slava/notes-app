package handler

import (
	_ "github.com/Tsygankov-Slava/notes-app/docs"

	"github.com/Tsygankov-Slava/notes-app/pkg/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		notes := api.Group("/notes")
		{
			notes.POST("/", h.createNote)
			notes.GET("/", h.getAllNotes)
			notes.GET("/:id", h.getNoteById)
			// TODO: notes.GET("/:id", h.getNotesByUserId)
			notes.PUT("/:id", h.updateNoteById)
			notes.DELETE("/:id", h.deleteNoteById)
			// TODO: notes.DELETE("/", h.deleteNotes)
		}
	}

	return router
}
