package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email    string `gorm:"unique;not null" json:"email"`
    Password string `json:"-"`
    Role     string `json:"role" gorm:"default:user"`
}
