package models


type Role struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"uniqueIndex;not null"`
    Description string
    Users       []User `gorm:"many2many:user_roles;"`
}
