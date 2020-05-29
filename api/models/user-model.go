package models

import (
	"log"
	s "notification-app/api/settings"
	"time"
)

// User is a struct for users
type User struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Email     string    `json:"email" gorm:"size:255;not null"`
	Phone     string    `json:"phone_number" gorm:"size:20"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// Create will crete a user data in the database
func (user *User) Create() error {
	if err := s.DB.Create(&user).Error; err != nil {
		log.Println("Ошибка при создании клиента: ", err)
		return err
	}
	return nil
}
