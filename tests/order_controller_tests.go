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

func TestPlaceOrder(t *testing.T) {
    reqBody := `{
        "user_id": 1,
        "items": [{"product_id": 1, "quantity": 2, "price": 10.5}]
    }`
    req, _ := http.NewRequest("POST", "/orders", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Request = req

    // Call the PlaceOrder function
    controllers.PlaceOrder(c)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "order placed successfully", response["message"])
}

func TestCancelOrder(t *testing.T) {
    req, _ := http.NewRequest("PUT", "/orders/1/cancel", nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Request = req

    // Call the CancelOrder function
    controllers.CancelOrder(c)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "order cancelled successfully", response["message"])
}