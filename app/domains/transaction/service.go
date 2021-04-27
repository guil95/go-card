package transaction

import (
	"errors"
	"log"

	entities "github.com/guil95/go-card/app/entities/transaction"
	"github.com/guil95/go-card/app/vo/uuid"

	"github.com/guil95/go-card/app/domains/account"

	"github.com/guil95/go-card/infra/repositories"
)

type Service struct {
	repo           *repositories.TransactionRepo
	accountService *account.Service
}

func NewService(r *repositories.TransactionRepo, as *account.Service) *Service {
	return &Service{repo: r, accountService: as}
}

func (s Service) MakeTransaction(accountID uuid.ID, amount float64, operationType entities.OperationType) (*entities.Transaction, error) {
	if !isValidOperationType(operationType) {
		log.Println("Invalid Operation type with value:", operationType)
		return nil, errors.New("Invalid Operation type")
	}

	account, err := s.accountService.FindAccountByID(accountID)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	transaction := entities.NewTransaction(accountID, operationType, amount)

	if transactionIsDebit(operationType) {
		transaction.Amount = amount * -1
		//TODO Atualizar account
	}

	if account.AvailableCreditLimit+transaction.Amount < 0 {
		return nil, entities.ErrorAccountCreditLimit
	}

	account.AvailableCreditLimit = account.AvailableCreditLimit + transaction.Amount

	_, err = s.repo.SaveTransaction(transaction)

	if err != nil {
		log.Println(err.Error)
		log.Println(account)
		return nil, err
	}

	_, err = s.accountService.UpdateAvailableCreditLimit(account)

	if err != nil {
		log.Println(err.Error)
		log.Println(account)
		return nil, err
	}

	return transaction, nil
}

func transactionIsDebit(operationType entities.OperationType) bool {
	for i := range entities.DebitTypes() {
		if entities.DebitTypes()[i] == operationType {
			return true
		}
	}

	return false
}

func isValidOperationType(operationType entities.OperationType) bool {
	for i := range entities.AllOperationsTypes() {
		if entities.AllOperationsTypes()[i] == operationType {
			return true
		}
	}

	return false
}
