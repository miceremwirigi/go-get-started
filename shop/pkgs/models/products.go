package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code         string `json:"code"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	Quantity     int    `json:"quantity"`
	HirePurchase bool   `json:"hire_purchase"`
}
