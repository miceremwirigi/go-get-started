package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNumber  string `json:"order_number"`
	ItemsQuantity int `json:"items_quantity"`
}

type OrderItem struct {
	gorm.Model
	CartItemID int	`json:"cart_item_id"`
	CartItem CartItem `json:"cart_item"`

	OrderID int `json:"order_id"`
	Order Order `json:"order"`

	Quantity int `json:"quantity"`
}