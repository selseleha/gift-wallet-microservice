package repositories

import (
	"task/pkg"
	"task/pkg/models"
)

type TransactionRepository interface {
	GetTransactions() ([]models.Transaction, error)
}

type TransactionRepositoryImpl struct {
	db *pkg.Database
}

func NewTransactionRepositoryImpl(db *pkg.Database) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (tr TransactionRepositoryImpl) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := tr.db.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
