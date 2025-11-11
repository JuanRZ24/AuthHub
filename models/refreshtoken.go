package models 

import ("time")

type RefreshToken struct {
    ID         uint      `gorm:"primaryKey"`
    SessionID  uint      `gorm:"index;not null"`
    TokenHash  string    `gorm:"not null"`
    ExpiresAt  time.Time `gorm:"not null"`
    RevokedAt  *time.Time
}
