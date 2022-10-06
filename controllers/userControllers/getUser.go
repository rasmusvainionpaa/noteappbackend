package usercontrollers

import (
	"net/http"
	"noteappbackend/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user, userExist := c.Get("user")

	if userExist == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "can't find user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        user.(models.User).ID,
		"firstName": user.(models.User).FirstName,
		"lastName":  user.(models.User).LastName,
		"email":     user.(models.User).Email,
	})
}
