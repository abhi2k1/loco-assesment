package datastore

import "github.com/loco-assessment/models"

type Repository interface {
	AddTransaction(txn models.Transaction) (err error)
	UpdateTransaction(txn models.Transaction) (err error)
	GetTxn(txnId int) (txn models.Transaction, err error)
	GetTxnsForEvent(txnType string) (txnIDs []int, err error)
	GetTxnsSum(txnId int) (sum float64, err error)
}
