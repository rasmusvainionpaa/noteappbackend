package routes

import (
	usercontrollers "noteappbackend/controllers/userControllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(gin *gin.Engine, db *gorm.DB) {
	handler := &usercontrollers.UserHandler{
		DB: db,
	}

	routes := gin.Group("/user")
	routes.GET("/", handler.GetUser)
	routes.GET("/:id", handler.GetUser)
	routes.PUT("/:id", handler.UpdateUser)
	routes.POST("/", handler.CreateUser)
	routes.DELETE("/:id", handler.DeleteUser)
}
