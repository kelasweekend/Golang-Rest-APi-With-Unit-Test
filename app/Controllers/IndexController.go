package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index ... Get all users
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Welcome To Home Page",
	})
}
