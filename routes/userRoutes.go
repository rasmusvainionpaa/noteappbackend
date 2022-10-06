package routes

import (
	usercontrollers "noteappbackend/controllers/userControllers"
	"noteappbackend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(gin *gin.Engine) {

	routes := gin.Group("/user")
	routes.GET("/", middlewares.RequireAuth, usercontrollers.GetUser)
	routes.PUT("/:id", middlewares.RequireAuth, usercontrollers.UpdateUser)
	routes.POST("/signup", usercontrollers.CreateUser)
	routes.POST("/login", usercontrollers.LoginUser)
	routes.DELETE("/:id", middlewares.RequireAuth, usercontrollers.DeleteUser)
}
