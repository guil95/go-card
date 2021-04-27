package repositories

import (
	"database/sql"

	entities "github.com/guil95/go-card/app/entities/account"
	"github.com/guil95/go-card/app/vo/uuid"
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

func (accountRepo *AccountRepo) FindAccountByDocument(document string) (*entities.Account, error) {
	stmt, err := accountRepo.db.Prepare(`SELECT id, document_number FROM accounts WHERE document_number = ?`)

	if err != nil {
		return nil, err
	}

	var account entities.Account

	rows, err := stmt.Query(document)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&account.ID, &account.Document)
	}

	if account.Document == "" {
		return nil, nil
	}

	return &account, nil
}

func (accountRepo *AccountRepo) FindAccountByID(id uuid.ID) (*entities.Account, error) {
	stmt, err := accountRepo.db.Prepare(`SELECT id, document_number FROM accounts WHERE id = ?`)

	if err != nil {
		return nil, err
	}

	var account entities.Account

	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&account.ID, &account.Document)
	}

	if account.Document == "" {
		return nil, nil
	}

	return &account, nil
}

func (accountRepo *AccountRepo) CreateAccount(account *entities.Account) (*entities.Account, error) {
	stmt, err := accountRepo.db.Prepare(`INSERT INTO accounts (id, document_number) values (?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(account.ID, account.Document)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return account, nil
}
