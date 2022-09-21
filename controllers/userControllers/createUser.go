package usercontrollers

import (
	"net/http"
	"noteappbackend/database"
	"noteappbackend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type requestBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
}

// handle singup
func CreateUser(c *gin.Context) {
	body := requestBody{}

	// read user from request
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error with password",
		})
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash),
	}

	// create user to database
	result := database.DB.Create((&user))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating user",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{})

}
