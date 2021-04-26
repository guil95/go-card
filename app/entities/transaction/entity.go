package entities

import (
	"time"

	"github.com/guil95/go-card/app/vo/uuid"
)

type Transaction struct {
	ID              uuid.ID       `json:"id"`
	AccountID       uuid.ID       `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
	EventDate       time.Time     `json:"event_date"`
}

func NewTransaction(accountID uuid.ID, operationTypeID OperationType, amount float64) *Transaction {
	return &Transaction{
		ID:              uuid.NewID(),
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       time.Now(),
	}
}
