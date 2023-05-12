package models

import (
	"gorm.io/gorm"
)

type Order struct {
    gorm.Model
    OrderNumber string       `json:"order_number" gorm:"primaryKey"`
    Products    string       `json:"products"`
    OrderItems  []*OrderItem `json:"order_items" gorm:"foreignKey:OrderNumber;references:OrderNumber"`
}

type OrderItem struct {
    gorm.Model
    CartNumber        int    `json:"cart_number"`
    Cart              *Cart  `json:"cart" gorm:"foreignKey:CartNumber;references:CartNumber"`

    OrderNumber       string `json:"order_number"`
    Order             *Order `json:"order" gorm:"foreignKey:OrderNumber;references:OrderNumber"`
    Quantity          int    `json:"quantity"`
}


