package repositories

import (
	"database/sql"

	entities "github.com/guil95/go-card/app/entities/account"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (accountRepo *AccountRepo) List() ([]*entities.Account, error) {

	stmt, err := accountRepo.db.Prepare(`select id, document_number from accounts`)
	if err != nil {
		return nil, err
	}
	var accounts []*entities.Account

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var a entities.Account

		err = rows.Scan(&a.ID, &a.Document)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &a)
	}

	return accounts, nil
}
