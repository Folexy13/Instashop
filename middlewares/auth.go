package middlewares

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "api/ecommerce/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization required"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(strings.TrimPrefix(tokenStr, "Bearer "), func(token *jwt.Token) (interface{}, error) {
            return utils.JwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }
        c.Next()
    }
}
