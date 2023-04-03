package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Customer   string     `gorm:"customer"`
	CartNumber string     `json:"cart_number"`
	CartItems  []CartItem `json:"cart_items"`
	// OrderItems []OrderItem `json:"order_items"`
}

type CartItem struct {
	gorm.Model
	ItemName  string  `json:"item_name"`
	ProductID int     `json:"product_id" gorm:"foreignKey:ID"`
	Product   Product `gorm:"foreignKey:ProductID;references:id" json:"product"`
	CartID    int     `gorm:"foreignKey:ID" json:"cart_id"`
	Cart      Cart    `gorm:"foreignKey:CartID;references:id" json:"cart"`
	Quantity  int     `json:"quantity"`
}
