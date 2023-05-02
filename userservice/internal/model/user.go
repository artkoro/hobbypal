package model

import (
    "time"

    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Name     string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    // Add other attributes as needed
    CreatedAt time.Time
    UpdatedAt time.Time
    // Add any other fields required
}