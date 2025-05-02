package routes

import (
	"merchant-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/merchant", controller.CreateMerchant)
	router.POST("/merchants/:id/topup", controller.TopUpBalance)
	router.POST("merchants/:id/withdraw", controller.WithdrawBalance)

	router.GET("/merchants", controller.GetMerchants)
	router.GET("/merchants/:id", controller.GetMerchantByID)
	router.GET("/merchants/:id/balance", controller.GetBalanceByMerchantID)

}
