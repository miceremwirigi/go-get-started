package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CartNumber string `json:"cart_number"`
	Quantity   int    `json:"quantity"`
}

type CartItem struct {
	gorm.Model
	ProductID int     `json:"product_id"`
	Product   Product `json:"product"`

	CartID int  `json:"cart_id"`
	Cart   Cart `json:"cart"`

	Quantity int `json:"quantity"`
}
