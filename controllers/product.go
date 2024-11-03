package controllers

import (
    "net/http"
    "api/ecommerce/config"
    "api/ecommerce/models"
    "github.com/gin-gonic/gin"
)

// Create a new product (admin only)
func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product := models.Product{Name: input.Name, Description: input.Description, Price: input.Price, Quantity: input.Quantity}
    config.DB.Create(&product)

    c.JSON(http.StatusCreated, gin.H{"product": product})
}

// Read all products
func GetProducts(c *gin.Context) {
    var products []models.Product
    config.DB.Find(&products)

    c.JSON(http.StatusOK, gin.H{"products": products})
}

// Update a product (admin only)
func UpdateProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Save(&product)
    c.JSON(http.StatusOK, gin.H{"product": product})
}

// Delete a product (admin only)
func DeleteProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    config.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
