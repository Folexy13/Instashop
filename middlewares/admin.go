package middlewares

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v4"
    "github.com/gin-gonic/gin"
    "api/ecommerce/utils"
)

func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the JWT token from the Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        // Parse and validate the JWT token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Validate the signing method and provide the secret key
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorSignatureInvalid)
            }
            return utils.JwtKey, nil 
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Check if the "isAdmin" claim is true
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !claims["isAdmin"].(bool) {
            c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            c.Abort()
            return
        }

        // Allow access if admin check passes
        c.Next()
    }
}
