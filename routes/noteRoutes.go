package routes

import (
	notecontrollers "noteappbackend/controllers/noteControllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterNoteRoutes(gin *gin.Engine, db *gorm.DB) {
	handler := &notecontrollers.NoteHandler{
		DB: db,
	}

	routes := gin.Group("/notes")
	routes.GET("/", handler.GetNotes)
	routes.GET("/:id", handler.GetNote)
	routes.PUT("/:id", handler.UpdateNote)
	routes.POST("/", handler.PostNote)
	routes.DELETE("/:id", handler.DeleteNote)
}
