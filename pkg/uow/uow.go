package uow

import (
	"context"
	"database/sql"

	applicationerrors "github.com/IcaroSilvaFK/picpay/application/errors"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Do(ctx context.Context, fc func(uow *Uow) *applicationerrors.ApplicationError) *applicationerrors.ApplicationError
	Rollback() *applicationerrors.ApplicationError
	CommitOrRollback() *applicationerrors.ApplicationError
	Register(name string, repository RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	UnRegister(name string)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) UowInterface {

	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, repository RepositoryFactory) {
	u.Repositories[name] = repository
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {

	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)

		if err != nil {
			return nil, err
		}

		u.Tx = tx
	}

	repo := u.Repositories[name](u.Tx)

	return repo, nil
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) Do(ctx context.Context, fc func(uow *Uow) *applicationerrors.ApplicationError) *applicationerrors.ApplicationError {

	if u.Tx != nil {
		return applicationerrors.InternalServerException()
	}

	tx, err := u.Db.BeginTx(ctx, nil)

	if err != nil {
		return applicationerrors.InternalServerException()
	}

	u.Tx = tx

	appErr := fc(u)

	if appErr != nil {
		errRb := u.Rollback()

		if errRb != nil {
			return applicationerrors.InternalServerException()
		}

		return appErr
	}

	return u.CommitOrRollback()
}

func (u *Uow) Rollback() *applicationerrors.ApplicationError {
	if u.Tx == nil {
		return applicationerrors.InternalServerException()
	}
	err := u.Tx.Rollback()

	if err != nil {
		return applicationerrors.InternalServerException()
	}

	u.Tx = nil
	return nil
}

func (u *Uow) CommitOrRollback() *applicationerrors.ApplicationError {

	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return applicationerrors.InternalServerException()
		}
		return applicationerrors.InternalServerException()
	}

	u.Tx = nil

	return nil

}
