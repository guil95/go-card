package entities

import (
	"errors"

	"github.com/guil95/go-card/app/vo/uuid"
)

var ErrorAccountNotFound = errors.New("account Not Found")

type Account struct {
	ID                   uuid.ID `json:"id"`
	Document             string  `json:"document_number"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}

func NewAccount(documentNumber string, availableCreditLimit float64) *Account {
	return &Account{
		ID:                   uuid.NewID(),
		Document:             documentNumber,
		AvailableCreditLimit: availableCreditLimit,
	}
}
