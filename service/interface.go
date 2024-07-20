package service

import "github.com/loco-assessment/models"

type Service interface {
	CreateTransaction(transactionId int, amount float64, transactionType string, parentId *int) error
	GetTransaction(transactionId int) (models.Transaction, error)
	GetAllTransactionEvent(transactionEvent string) ([]int, error)
	GetTransactionSum(transactionId int) (float64, error)
}
