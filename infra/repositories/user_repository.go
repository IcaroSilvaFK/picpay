package repositories

import (
	"github.com/IcaroSilvaFK/picpay/infra/database"
	"github.com/IcaroSilvaFK/picpay/infra/models"
)

type UserRepository struct {
	db database.Queryer
}

type UserRepositoryInterface interface {
	Create(u *models.UserModel) error
	GetById(id int) (*models.UserModel, error)
	Delete(id int) error
}

func NewUserRepository(db database.Queryer) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(u *models.UserModel) error {
	stmt, err := ur.db.Prepare(`INSERT INTO users (name,email,password,identifier, type) VALUES ($1,$2,$3, $4, $5)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.Password, u.Identifier, u.Type)

	return err
}

func (ur *UserRepository) GetById(id int) (*models.UserModel, error) {

	stmt, err := ur.db.Prepare(`SELECT * FROM users WHERE id = $1`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var u models.UserModel

	row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Identifier,
		&u.Type,
	)

	return &u, nil
}

func (ur *UserRepository) Delete(id int) error {

	stmt, err := ur.db.Prepare(`DELETE FROM users WHERE id = $1`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err

}
