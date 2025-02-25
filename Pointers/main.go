package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?

func updateBalance(cus *customer, tran transaction) error {
	if tran.transactionType == transactionDeposit {
		cus.balance += tran.amount
	} else if tran.transactionType == transactionWithdrawal {
		if tran.amount > cus.balance {
			return errors.New("insufficient funds")
		}
		cus.balance -= tran.amount
	} else {
		return errors.New("unknown transaction type")
	}

	return nil
}
