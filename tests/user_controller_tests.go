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

func TestRegisterUser(t *testing.T) {
    

    // Simulate request body
    reqBody := `{"email": "test@example.com", "password": "password"}`
    req, _ := http.NewRequest("POST", "/register", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    
    // Create a response recorder
    w := httptest.NewRecorder()

    // Initialize a new Gin context for testing
    c, _ := gin.CreateTestContext(w)
    c.Request = req

    // Call the Register function, passing in the test context
    controllers.Register(c)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.NotEmpty(t, response["token"])
}

func TestLoginUser(t *testing.T) {
    

    // Simulate request body
    reqBody := `{"email": "test@example.com", "password": "password"}`
    req, _ := http.NewRequest("POST", "/login", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")

    // Create a response recorder
    w := httptest.NewRecorder()

    // Initialize a new Gin context for testing
    c, _ := gin.CreateTestContext(w)
    c.Request = req

    // Call the Login function, passing in the test context
    controllers.Login(c)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.NotEmpty(t, response["token"])
}