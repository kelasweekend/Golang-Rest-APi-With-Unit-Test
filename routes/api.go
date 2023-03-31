// Routes/Routes.go
package api

import (
	"go-restapi-mysql/app/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	route := gin.Default()

	route.GET("/", Controllers.GetIndex)
	activity := route.Group("/api/activity")
	{
		activity.GET("", Controllers.GetActivity)
		// activity.POST("user", Controllers.CreateUser)
		// activity.GET("user/:id", Controllers.GetUserByID)
		// activity.PUT("user/:id", Controllers.UpdateUser)
		// activity.DELETE("user/:id", Controllers.DeleteUser)
	}
	return route
}
