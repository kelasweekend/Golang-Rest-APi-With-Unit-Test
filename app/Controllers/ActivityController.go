package Controllers

import (
	"go-restapi-mysql/app/Models/Activity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index ... Get all Activity
func GetActivity(c *gin.Context) {
	var activity []Activity.Activity
	err := Activity.GetAll(&activity)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Get All Activity",
		"data":    activity,
	})
}
