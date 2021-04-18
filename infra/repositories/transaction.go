package repositories

import (
	"database/sql"

	entities "github.com/guil95/go-card/app/entities/transaction"
)

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

func (transactionRepo *TransactionRepo) SaveTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	stmt, err := transactionRepo.db.Prepare(`
	INSERT INTO transactions 
	(transaction_uuid, account_id, operation_type_id, amount, event_date) 
	values (?,?,?,?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(transaction.ID, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
