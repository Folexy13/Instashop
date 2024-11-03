package controllers

import (
    "net/http"
    "api/ecommerce/config"
    "api/ecommerce/models"
    "github.com/gin-gonic/gin"
)

// Place an order for one or more products
func PlaceOrder(c *gin.Context) {
    var orderInput struct {
        UserID uint                `json:"user_id"`
        Items  []models.OrderItem  `json:"items"`
    }
    if err := c.ShouldBindJSON(&orderInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var totalPrice float64
    for _, item := range orderInput.Items {
        totalPrice += item.Price * float64(item.Quantity)
    }

    order := models.Order{
        UserID:     orderInput.UserID,
        Status:     "pending",
        TotalPrice: totalPrice,
        Items:      orderInput.Items,
    }

    if err := config.DB.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "order placed successfully", "order": order})
}

// List all orders for a specific user
func GetUserOrders(c *gin.Context) {
    userID := c.Param("user_id")

    var orders []models.Order
    if err := config.DB.Where("user_id = ?", userID).Preload("Items").Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}

// Cancel an order if it is still in Pending status
func CancelOrder(c *gin.Context) {
    orderID := c.Param("id")

    var order models.Order
    if err := config.DB.First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
        return
    }

    if order.Status != "pending" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "only pending orders can be cancelled"})
        return
    }

    order.Status = "cancelled"
    if err := config.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to cancel order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "order cancelled successfully"})
}

// Update the status of an order (admin privilege)
func UpdateOrderStatus(c *gin.Context) {
    orderID := c.Param("id")

    var input struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var order models.Order
    if err := config.DB.First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
        return
    }

    order.Status = input.Status
    if err := config.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "order status updated successfully"})
}
