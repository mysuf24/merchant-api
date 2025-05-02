package controller

import (
	"log"
	"merchant-api/config"
	"merchant-api/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateMerchant(ctx *gin.Context) {
	var input model.Merchant

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	//save merchant
	if err := config.DB.Create(&input).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membuat merchant"})
		return
	}

	log.Printf("Merchant berhasil dibuat: %+v\n", input)

	//setelah berhasil, buat saldo awal
	balance := model.MerchantBalance{
		MerchantID:       input.ID,
		AvailableBalance: 1000,
		HoldBalance:      0,
		UpdatedAt:        time.Now(),
	}

	if err := config.DB.Create(&balance).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memasukan saldo awal"})
		return
	}

	log.Printf("Saldo awal berhasil dibuat untuk merchant %s\n", input.ID)

	ctx.JSON(http.StatusCreated, gin.H{
		"merchant_id": input.ID,
		"status":      "success",
		"message":     "Merchant registered successfully",
	})

}

// get merchant
func GetMerchants(ctx *gin.Context) {
	var merchants []model.Merchant
	config.DB.Find(&merchants)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List Data merch",
		"data":    merchants,
	})
}

// get merchant by id
func GetMerchantByID(ctx *gin.Context) {
	var merchants model.Merchant
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&merchants).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail Data By Id :" + ctx.Param("id"),
		"data":    merchants,
	})
}


