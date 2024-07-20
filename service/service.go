package service

import (
	"errors"
	"github.com/loco-assessment/datastore"
	"github.com/loco-assessment/models"
	"strings"
)

type service struct {
	repo datastore.Repository
}

func NewService(repo datastore.Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateTransaction(transactionId int, amount float64, transactionType string, parentId *int) error {
	txn := models.Transaction{
		ID:                  transactionId,
		Amount:              amount,
		TransactionType:     strings.ToLower(transactionType),
		ParentTransactionID: parentId,
	}

	if parentId != nil && *parentId == transactionId {
		return errors.New("parent id cannot be same as transaction id")
	}

	return s.repo.AddTransaction(txn)
}

func (s *service) GetTransaction(transactionId int) (models.Transaction, error) {
	return s.repo.GetTxn(transactionId)
}

func (s *service) GetAllTransactionEvent(transactionEvent string) ([]int, error) {
	return s.repo.GetTxnsForEvent(strings.ToLower(transactionEvent))
}

func (s *service) GetTransactionSum(transactionId int) (float64, error) {
	return s.repo.GetTxnsSum(transactionId)
}
