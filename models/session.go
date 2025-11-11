package models

import ("time")

type Session struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"index;not null"`
    Device      string
    IP          string
    IsRevoked   bool      `gorm:"default:false"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    RefreshToken RefreshToken
}
