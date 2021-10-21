package models

type Wallet struct {
	Id          int32  `gorm:"primary_key;AUTO_INCREMENT"`
	PhoneNumber string `gorm:"not null"`
	Amount      int32
}

func (Wallet) TableName() string {
	return "wallet"
}
