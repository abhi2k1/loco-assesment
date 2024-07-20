package inmemory

import (
	"errors"
	"fmt"
	"github.com/loco-assessment/models"
	"sync"
)

type inMemoryDatastore struct {
	transactionMap map[int]models.Transaction // map of transaction ID to transaction
	parentMap      map[int][]int              // map of parent transaction ID to child transaction IDs
	mtx            *sync.Mutex                // mutex to protect the maps
}

func NewInMemoryDatastore() *inMemoryDatastore {
	return &inMemoryDatastore{
		transactionMap: make(map[int]models.Transaction),
		parentMap:      make(map[int][]int),
		mtx:            &sync.Mutex{},
	}
}

func (i *inMemoryDatastore) AddTransaction(txn models.Transaction) (err error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	_, ok := i.transactionMap[txn.ID]
	if ok {
		return errors.New("Transaction already exist")
	}

	i.transactionMap[txn.ID] = txn

	if txn.ParentTransactionID != nil {
		i.parentMap[*txn.ParentTransactionID] = append(i.parentMap[*txn.ParentTransactionID], txn.ID)
	}

	return nil
}

func (i *inMemoryDatastore) UpdateTransaction(txn models.Transaction) (err error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	oldTxn, ok := i.transactionMap[txn.ID]
	if !ok {
		return errors.New("txn not found")
	}

	if oldTxn.ParentTransactionID != nil {
		parentTxns, ok := i.parentMap[*oldTxn.ParentTransactionID]
		if ok {
			for j := range parentTxns {
				if parentTxns[j] == txn.ID {
					parentTxns = append(parentTxns[:j], parentTxns[j+1:]...)
					break
				}
			}
		}
	}

	fmt.Println(oldTxn)

	i.transactionMap[txn.ID] = txn
	if txn.ParentTransactionID != nil {
		i.parentMap[*txn.ParentTransactionID] = append(i.parentMap[*txn.ParentTransactionID], txn.ID)
	}

	return nil
}

func (i *inMemoryDatastore) GetTxn(txnId int) (txn models.Transaction, err error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	txn, ok := i.transactionMap[txnId]
	if !ok {
		err = errors.New("txn not found")
		return
	}

	return txn, nil
}

func (i *inMemoryDatastore) GetTxnsForEvent(txnType string) (txnIDs []int, err error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	txnIDs = make([]int, 0)

	for _, txn := range i.transactionMap {
		if txn.TransactionType == txnType {
			txnIDs = append(txnIDs, txn.ID)
		}
	}

	return txnIDs, nil
}

func (i *inMemoryDatastore) GetTxnsSum(txnId int) (sum float64, err error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	txn, ok := i.transactionMap[txnId]
	if !ok {
		err = errors.New("txn not found")
		return
	}

	sum += txn.Amount

	temp, err := i.getAllChildSums(txnId)
	if err != nil {
		return 0, err
	}

	sum += temp

	return sum, nil

}

func (i *inMemoryDatastore) getAllChildSums(txnId int) (float64, error) {
	if txnId == 0 {
		return 0, nil
	}

	txns, ok := i.parentMap[txnId]
	if !ok {
		return 0, nil
	}

	sum := 0.0

	for j := range txns {
		data, ok := i.transactionMap[txns[j]]
		if !ok {
			return 0, errors.New("txn not found")
		}

		temp, err := i.getAllChildSums(data.ID)
		if err != nil {
			return 0, err
		}

		fmt.Println("temp", temp, data.Amount, "id", data.ID)

		sum += temp + data.Amount
	}

	return sum, nil
}
