package transaction

import (
	"errors"
	"log"

	"github.com/guil95/go-card/app/domains/account"

	"github.com/guil95/go-card/app/utils"
	"github.com/guil95/go-card/infra/repositories"
)

type Service struct {
	repo           *repositories.TransactionRepo
	accountService *account.Service
}

func NewService(r *repositories.TransactionRepo, as *account.Service) *Service {
	return &Service{repo: r, accountService: as}
}

func (s Service) MakeTransaction(accountID utils.ID, amount float64, operationType OperationType) (bool, error) {
	if !isValidOperationType(operationType) {
		log.Println("Invalid Operation type with value:", operationType)
		return false, errors.New("Invalid Operation type")
	}

	account, err := s.accountService.FindAccountByID(accountID)

	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	amount = retrieveAmountType(operationType, amount)

	return false, nil
}

func retrieveAmountType(operationType OperationType, amount float64) (calculatedAmount float64) {
	for i := range DebitTypes() {
		if DebitTypes()[i] == operationType {
			return amount * -1
		}
	}

	return amount
}

func isValidOperationType(operationType OperationType) bool {
	for i := range AllOperationsTypes() {
		if AllOperationsTypes()[i] == operationType {
			return true
		}
	}

	return false
}
