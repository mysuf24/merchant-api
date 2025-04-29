package model

import "time"

type Merchant struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	LegalName    string    `json:"legal_name"`
	MerchantType string    `json:"merchant_type"`
	NPWPNumber   string    `json:"npwp_number"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	Province     string    `json:"province"`
	PostalCode   string    `json:"postal_code"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Balances []MerchantBalance `gorm:"foreignKey:MerchantID"`
}

type MerchantBalance struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	MerchantID       string    `gorm:"index" json:"merchant_id"`
	Merchant         Merchant  `gorm:"foreignKey:MerchantID;references:ID"`
	AvailableBalance int64     `json:"available_balance"`
	HoldBalance      int64     `json:"hold_balance"`
	UpdatedAt        time.Time `json:"updated_at"`
}
