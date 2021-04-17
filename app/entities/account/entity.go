package entities

import "github.com/google/uuid"

type Account struct {
	ID       uuid.UUID `json:"id"`
	Document string    `json:"document_number"`
}
