package tests

import (
    "api/ecommerce/config"
    "api/ecommerce/models"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "os"
    "testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
    os.Setenv("GIN_MODE", "test")

    // Initialize test database (SQLite in-memory for simplicity)
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrate the schema
    db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{})

    config.DB = db // Set the global DB to the test DB

    router = gin.Default()
    // Register routes for testing purposes
    // routes.RegisterRoutes(router)

    // Run tests
    code := m.Run()
    os.Exit(code)
}