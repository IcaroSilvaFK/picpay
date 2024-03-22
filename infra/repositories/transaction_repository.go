package repositories

import (
	"github.com/IcaroSilvaFK/picpay/infra/database"
	"github.com/IcaroSilvaFK/picpay/infra/models"
)

type TransactionRepository struct {
	db database.Queryer
}

type TransactionRepositoryInterface interface {
	Create(*models.TransactionModel) error
	GetMyTransactions(userId int) ([]*models.TransactionModel, error)
	Delete(id string) error
}

func NewTransactionRepository(db database.Queryer) TransactionRepositoryInterface {
	return &TransactionRepository{
		db: db,
	}
}

func (tr *TransactionRepository) Create(tm *models.TransactionModel) error {

	stmt, err := tr.db.Prepare(`INSERT INTO transactions (payer, payee, value) VALUES ($1,$2,$3)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(tm.Payer, tm.Payee, tm.Amount)

	return err
}

func (tr *TransactionRepository) GetMyTransactions(userId int) ([]*models.TransactionModel, error) {

	stmt, err := tr.db.Prepare(`SELECT * FROM transactions WHERE payer = $1 OR payee = $2`)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId, userId)

	if err != nil {
		return nil, err
	}

	var r []*models.TransactionModel

	for rows.Next() {

		var t models.TransactionModel

		rows.Scan(
			t.ID,
			t.Payer,
			t.Payee,
			t.Amount,
		)

		r = append(r, &t)
	}

	return r, nil
}

func (tr *TransactionRepository) Delete(id string) error {

	stmt, err := tr.db.Prepare(`DELETE FROM transactions WHERE id = $1`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}
