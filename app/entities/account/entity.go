package entities

import (
	"github.com/google/uuid"
	"github.com/guil95/go-card/app/utils"
)

type Account struct {
	ID       uuid.UUID `json:"id"`
	Document string    `json:"document_number"`
}

func NewAccount(documentNumber string) *Account {
	return &Account{
		ID:       utils.NewID(),
		Document: documentNumber,
	}
}
