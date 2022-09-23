package notecontrollers

import (
	"net/http"
	"noteappbackend/models"

	"github.com/gin-gonic/gin"
)

func GetNote(c *gin.Context) {
	user, userExist := c.Get("user")

	if userExist == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "can't find user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.(models.User).ID,
	})
}
