package config

import (
	"log"
	"merchant-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // variable untuk diakses dari controller

func ConnectDB() {
	dns := "host=localhost user=postgres password=123 port=5432 dbname=db_merchant sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("gagal untuk connect ke database")
	}

	database.AutoMigrate(&model.Merchant{}, &model.MerchantBalance{})

	DB = database
}
