package routes

import (
	"merchant-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/merchant", controller.CreateMerchant)
	router.GET("/merchants", controller.GetMerchants)
	router.GET("/merchants/:id", controller.GetMerchantByID)

}
