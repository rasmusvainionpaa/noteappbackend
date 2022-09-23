package notecontrollers

import (
	"net/http"
	"noteappbackend/database"
	"noteappbackend/models"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	var notes []models.Note

	user, userExist := c.Get("user")

	if userExist == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "can't find user",
		})
		return
	}

	if result := database.DB.Find(&notes, "user_id = ?", user.(models.User).ID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &notes)
}
