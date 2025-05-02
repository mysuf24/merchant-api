package controller

import (
	"log"
	"merchant-api/config"
	"merchant-api/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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
	if err := config.DB.Clauses(clause.Returning{}).Create(&input).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membuat merchant"})
		return
	}

	log.Printf("Merchant berhasil dibuat: %+v\n", input.ID)

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

	if err := config.DB.Preload("Balances").Find(&merchants).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data merchant"})
		return
	}

	log.Printf("Berhasil ambil %d merchant\n", len(merchants))
	ctx.JSON(200, gin.H{
		"success": true,
		"message": "List Data merch",
		"data":    merchants,
	})
}

// get merchant by id
func GetMerchantByID(ctx *gin.Context) {
	var output model.Merchant

	if err := config.DB.Preload("Balances").Where("id = ?", ctx.Param("id")).First(&output).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	log.Printf("Merchant ID %s ditemukan\n", ctx.Param("id"))
	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data By Id :" + ctx.Param("id"),
		"data":    output,
	})
}

// get balance merchant by id
func GetBalanceByMerchantID(ctx *gin.Context) {
	merchantID := ctx.Param("id")

	var balance model.MerchantBalance
	if err := config.DB.Where("merchant_id = ?", merchantID).First(&balance).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Balance kosong pada merchant ini"})
		return
	}

	log.Printf("Saldo ditemukan untuk merchant_id %s\n", merchantID)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "saldo merchant di temukan",
		"data":    balance,
	})
}

func TopUpBalance(ctx *gin.Context) {
	merchantID := ctx.Param("id")
	var req struct {
		Amount int64 `json:"amount"`
	}

	if err := ctx.ShouldBind(&req); err != nil || req.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Jumlah top-up harus lebih dari 0"})
		return
	}

	var balance model.MerchantBalance
	if err := config.DB.Where("merchant_id = ?", merchantID).First(&balance).Error; err != nil {
		log.Printf("Saldo untuk merchant %s tidak ditemukan", merchantID)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Merchant tidak ditemukan atau belum punya saldo"})
		return
	}

	oldBalance := balance.AvailableBalance
	balance.AvailableBalance += req.Amount
	balance.UpdatedAt = time.Now()

	if err := config.DB.Save(&balance).Error; err != nil {
		log.Printf("Gagal update saldo untuk merchant  %s: %v", merchantID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui saldo"})
		return
	}

	log.Printf("Top-up berhasil untuk merchant %s, dari %d menjadi %d", merchantID, oldBalance, balance.AvailableBalance)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Top up berhasil",
		"data":    balance,
	})
}
