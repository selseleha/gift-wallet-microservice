package repositories

import (
	"task/pkg"
	"task/pkg/models"
	"time"
)

type TransactionRepository interface {
	CreateTransaction(phoneNumber string, amount int32, operation int32) error
}

type TransactionRepositoryImpl struct {
	db *pkg.Database
}

func NewTransactionRepositoryImpl(db *pkg.Database) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (tr TransactionRepositoryImpl) CreateTransaction(phoneNumber string, amount int32, operation int32) error {
	err := tr.db.DB.Create(&models.Transaction{
		Amount:      amount,
		Operation:   operation,
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
	}).Error
	return err
}
