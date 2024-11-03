package main

import (
	"github.com/gin-gonic/gin"
	"api/ecommerce/config"
	"api/ecommerce/routes"
)

func main() {
	config.InitDatabase()

	router := gin.Default()
	
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
