package models

import (
	"log"
	s "notification-app/api/settings"
	"time"
)

// Order is a struct for orders
type Order struct {
	ID         int       `json:"id" gorm:"primary_key;auto_increment"`
	Goods      string    `json:"goods"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

//Create will crete a order data in the database
func (order *Order) Create() error {
	if err := s.DB.Create(&order).Error; err != nil {
		log.Println("Ошибка при создании покупок: ", err)
		return err
	}
	return nil
}
