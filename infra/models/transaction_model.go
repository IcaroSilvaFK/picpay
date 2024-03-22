package models

type TransactionModel struct {
	ID     int
	Payer  int
	Payee  int
	Amount int
}

func NewTransactionModel(
	payer, payee, amount int,
) *TransactionModel {
	return &TransactionModel{
		Payer:  payer,
		Payee:  payee,
		Amount: amount,
	}
}
