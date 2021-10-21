package repositories

import (
	"task/pkg"
)

type WalletRepository interface {
}

type WalletRepositoryImpl struct {
	db *pkg.Database
}

func NewWalletRepositoryImpl(db *pkg.Database) *WalletRepositoryImpl {
	return &WalletRepositoryImpl{db: db}
}
