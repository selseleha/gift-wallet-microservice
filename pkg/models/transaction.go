package models

import "time"

type Transaction struct {
	Id          int32  `gorm:"primary_key;AUTO_INCREMENT"`
	Amount      int32  `gorm:"not null"`
	Operation   int32  `gorm:"not null"`
	PhoneNumber string `gorm:"not null"`
	CreatedAt   time.Time
}

func (Transaction) TableName() string {
	return "transaction"
}
