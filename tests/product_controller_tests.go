package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "api/ecommerce/controllers"
)

func TestCreateProduct(t *testing.T) {
    // Set up router and request body
    reqBody := `{"name": "Test Product", "price": 20.5}`
    req, _ := http.NewRequest("POST", "/products", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    // Create test context and assign request
    c, _ := gin.CreateTestContext(w)
    c.Request = req

    // Call the CreateProduct function
    controllers.CreateProduct(c)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "product created successfully", response["message"])
}

func TestGetProducts(t *testing.T) {
    req, _ := http.NewRequest("GET", "/api/products", nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Request = req

    // Call the GetProduct function
    controllers.GetProducts(c)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "Test Product", response["product"].(map[string]interface{})["name"])
}