package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code          string `json:"code"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	HirePurchase bool   `json:"hire_purchase"`
}
