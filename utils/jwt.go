package utils

import (
   "github.com/golang-jwt/jwt/v4"
    "time"
    "os"
    
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint, role string) (string, error) {
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
        IssuedAt:  time.Now().Unix(),
        Subject:   string(userID),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtKey)
}
