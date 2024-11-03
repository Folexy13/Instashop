package controllers

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "api/ecommerce/models"
    "api/ecommerce/utils"
    "api/ecommerce/config"
    "net/http"
)

func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    input.Password = string(hashedPassword)
    config.DB.Create(&input)
    c.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}

func Login(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
        return
    }

    // Check if the user role is admin and return isAdmin field accordingly
    isAdmin := user.Role == "admin"
    c.JSON(http.StatusOK, gin.H{
        "token":   token,
        "isAdmin": isAdmin,
    })
}