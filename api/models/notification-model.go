package models

import (
	"log"
	s "notification-app/api/settings"
	"time"
)

// Notification stuct to store data about notifications
type Notification struct {
	ID      int `json:"id" gorm:"primary_key;auto_increment"`
	UserID  int `json:"user_id" gorm:"not null"`
	OrderID int `json:"order_id" gorm:"not null"`

	SendByEmail *int      `json:"send_by_email" gorm:"default:0"`
	SendBySMS   *int      `json:"send_by_smsm" gorm:"default:0"`
	Message     string    `json:"message" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// Create will crete a notification data in the database
func (notification *Notification) Create() error {
	if err := s.DB.Create(&notification).Error; err != nil {
		log.Println("Ошибка при создании уведомления: ", err)
		return err
	}
	return nil
}
