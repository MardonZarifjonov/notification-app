package routers

import (
	"notification-app/api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter will set all APIs
func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/notification", controllers.SendNotification)
	}
	return router
}
