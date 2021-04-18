package transaction

type OperationType int64

const (
	PURCHASE_IN_CASH OperationType = iota + 1
	PURCHASE_INSTALLMENT
	WITHDRAW
	PAYMENT
)

func AllOperationsTypes() []OperationType {
	return []OperationType{
		PURCHASE_IN_CASH,
		PURCHASE_INSTALLMENT,
		WITHDRAW,
		PAYMENT,
	}
}

func DebitTypes() []OperationType {
	return []OperationType{
		PURCHASE_IN_CASH,
		PURCHASE_INSTALLMENT,
		WITHDRAW,
	}
}

func CreditTypes() []OperationType {
	return []OperationType{
		PAYMENT,
	}
}
