package utils

import "errors"

var (
	BadRequestError     = errors.New("bad request")
	WalletExistError    = errors.New("wallet is exist")
	WalletNouFoundError = errors.New("wallet not found")
)
