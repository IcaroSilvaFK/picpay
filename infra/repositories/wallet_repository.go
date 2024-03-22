package repositories

import (
	"github.com/IcaroSilvaFK/picpay/infra/database"
	"github.com/IcaroSilvaFK/picpay/infra/models"
)

type WalletRepository struct {
	db database.Queryer
}

type WalletRepositoryInterface interface {
	Create(wallet *models.WalletModel) error
	GetByUserId(userId int) (*models.WalletModel, error)
	Update(userId, balance int) error
	Delete(id int) error
}

func NewWalletRepository(db database.Queryer) WalletRepositoryInterface {
	return &WalletRepository{
		db: db,
	}
}

func (wr *WalletRepository) Create(wallet *models.WalletModel) error {
	stmt, err := wr.db.Prepare(`INSERT INTO wallets (user_id,amount) VALUES ($1, $2)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(wallet.UserId, wallet.Amount)

	return err
}

func (wr *WalletRepository) GetByUserId(userId int) (*models.WalletModel, error) {

	stmt, err := wr.db.Prepare(`SELECT * FROM wallets WHERE user_id = $1`)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(userId)

	var walletModel models.WalletModel

	if err := row.Scan(&walletModel.ID, &walletModel.UserId, &walletModel.Amount); err != nil {
		return nil, err
	}

	return &walletModel, nil
}

func (wr *WalletRepository) Update(userId, balance int) error {

	stmt, err := wr.db.Prepare(`UPDATE wallets SET amount = $1 WHERE user_id = $2`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(balance, userId)

	return err
}

func (wr *WalletRepository) Delete(id int) error {

	stmt, err := wr.db.Prepare(`DELETE FROM wallets WHERE id = $1`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}
