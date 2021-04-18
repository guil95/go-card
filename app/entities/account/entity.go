package entities

import (
	"errors"

	"github.com/guil95/go-card/app/utils"
)

var ErrorAccountNotFound = errors.New("Account Not Found")

type Account struct {
	ID       utils.ID `json:"id"`
	Document string   `json:"document_number"`
}

func NewAccount(documentNumber string) *Account {
	return &Account{
		ID:       utils.NewID(),
		Document: documentNumber,
	}
}
