package models

type Gift struct {
	Id          int32 `gorm:"primary_key;AUTO_INCREMENT"`
	PhoneNumber string
	Code        string `gorm:"not null"`
	Amount      int32
}

func (Gift) TableName() string {
	return "gift"
}
