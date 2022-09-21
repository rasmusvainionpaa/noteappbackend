package routes

import (
	notecontrollers "noteappbackend/controllers/noteControllers"
	"noteappbackend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterNoteRoutes(gin *gin.Engine) {

	routes := gin.Group("/notes")
	routes.GET("/", middlewares.RequireAuth, notecontrollers.GetNotes)
	routes.GET("/:id", middlewares.RequireAuth, notecontrollers.GetNote)
	routes.PUT("/:id", middlewares.RequireAuth, notecontrollers.UpdateNote)
	routes.POST("/", middlewares.RequireAuth, notecontrollers.PostNote)
	routes.DELETE("/:id", middlewares.RequireAuth, notecontrollers.DeleteNote)
}
