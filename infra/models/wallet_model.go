package models

type WalletModel struct {
	ID     int
	UserId int
	Amount int
}

func NewWalletModel(
	userId, amount int,
) *WalletModel {
	return &WalletModel{
		UserId: userId,
		Amount: amount,
	}
}
