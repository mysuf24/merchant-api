package routes

import (
	"merchant-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	//Tes untuk memastikan API aktif
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running"})
	})

	//untuk membuat merchant baru
	router.POST("/merchant", controller.CreateMerchant)
}
