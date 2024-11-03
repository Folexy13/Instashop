package routes

import (
    "github.com/gin-gonic/gin"
    "api/ecommerce/controllers"
    "api/ecommerce/middlewares"
    "net/http"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the instashop E-commerce API"})
}

func RegisterRoutes(router *gin.Engine) {
    
    router.GET("/", rootHandler)
    api := router.Group("/api")
    
    api.POST("/register", controllers.Register)
    api.POST("/login", controllers.Login)

    // Protected routes
    auth := api.Group("/")
    auth.Use(middlewares.AuthMiddleware())
    // Product routes (admin protected)
    adminRoutes := auth.Group("/products")
    adminRoutes.Use(middlewares.AdminOnly())
    {
        adminRoutes.POST("/", controllers.CreateProduct)
        adminRoutes.PUT("/:id", controllers.UpdateProduct)
        adminRoutes.DELETE("/:id", controllers.DeleteProduct)
    }
    
    
    api.GET("/products", controllers.GetProducts)
    
    // Order routes
    auth.POST("/orders", controllers.PlaceOrder)
    auth.GET("/orders/:user_id", controllers.GetUserOrders)
    adminRoutes.PUT("/orders/:id/cancel", controllers.CancelOrder)
}
