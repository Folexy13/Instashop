package models


import (
    "gorm.io/gorm"
)

type Order struct {
    gorm.Model
    UserID     uint         `json:"user_id"`
    Status     string       `json:"status" gorm:"default:pending"` // "pending", "completed", "cancelled"
    TotalPrice float64      `json:"total_price"`
    Items      []OrderItem  `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
    gorm.Model
    OrderID   uint    `json:"order_id"`
    ProductID uint    `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"` // Price at the time of purchase
}
