package main

import (
	"fmt"
	"net/http"
	"noteappbackend/database"
	"noteappbackend/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	// Connect to database
	database.Connect()

	// load env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Error loadin env file")
	}

	//add routes
	routes.RegisterUserRoutes(router)
	routes.RegisterNoteRoutes(router)

	// gets run when requested route isn't found
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": "404"})
	})

	//starts server on spesified port
	router.Run(os.Getenv("SERVER_PORT"))
}
