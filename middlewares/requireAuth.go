package middlewares

import (
	"fmt"
	"net/http"
	"noteappbackend/database"
	"noteappbackend/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// read tokenfrom request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read token",
		})
		return
	}

	// Check token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Token expired",
			})
			return
		}

		// find user with token id
		var user models.User

		database.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Coudn't find user",
			})
			return
		}

		// set user
		c.Set("user", user)

		// continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
