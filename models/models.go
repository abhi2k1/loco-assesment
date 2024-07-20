package models

type Transaction struct {
	ID                  int     `json:"transaction_id"`
	Amount              float64 `json:"amount"`
	TransactionType     string  `json:"transaction_type"`
	ParentTransactionID *int    `json:"parent_id"`
}
