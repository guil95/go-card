package entities

import (
	"errors"

	"github.com/guil95/go-card/app/vo/uuid"
)

var ErrorAccountNotFound = errors.New("Account Not Found")

type Account struct {
	ID       uuid.ID `json:"id"`
	Document string  `json:"document_number"`
}

func NewAccount(documentNumber string) *Account {
	return &Account{
		ID:       uuid.NewID(),
		Document: documentNumber,
	}
}
