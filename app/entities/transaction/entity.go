package entities

import (
	"time"

	"github.com/guil95/go-card/app/utils"
)

type Transaction struct {
	ID              utils.ID      `json:"id"`
	AccountID       utils.ID      `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
	EventDate       time.Time     `json:"event_date"`
}

func NewTransaction(accountID utils.ID, operationTypeID OperationType, amount float64) *Transaction {
	return &Transaction{
		ID:              utils.NewID(),
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       time.Now(),
	}
}
