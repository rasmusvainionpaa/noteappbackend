package notecontrollers

import (
	"net/http"
	"noteappbackend/database"
	"noteappbackend/models"

	"github.com/gin-gonic/gin"
)

type postBody struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Alert     string `json:"alert"`
	Important bool   `json:"isImportant"`
}

func PostNote(c *gin.Context) {
	user, userExist := c.Get("user")

	if userExist == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "can't find user",
		})
		return
	}

	body := postBody{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	note := models.Note{
		UserID:    float32(user.(models.User).ID),
		Title:     body.Title,
		Content:   body.Content,
		Alert:     body.Alert,
		Important: body.Important,
	}

	result := database.DB.Create(&note)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating note",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{})

}
