package services

import (
	"context"

	applicationerrors "github.com/IcaroSilvaFK/picpay/application/errors"
	"github.com/IcaroSilvaFK/picpay/infra/models"
	"github.com/IcaroSilvaFK/picpay/infra/repositories"
	"github.com/IcaroSilvaFK/picpay/pkg/uow"
)

type WalletService struct {
	uow uow.UowInterface
}

type WalletServiceInterface interface {
	Create(userId int, balance float64) *applicationerrors.ApplicationError
	UpdateBalance(userId int, balance float64) *applicationerrors.ApplicationError
	GetBalance(userId int) (*models.WalletModel, error)
	Delete(userId int) *applicationerrors.ApplicationError
}

func NewWalletService(uow uow.UowInterface) WalletServiceInterface {
	return &WalletService{uow}
}

func (ws *WalletService) Create(userId int, balance float64) *applicationerrors.ApplicationError {

	amount := int(balance * 1000)

	wallet := models.NewWalletModel(userId, amount)

	repo := ws.getRepository()

	err := repo.Create(wallet)

	if err != nil {
		return applicationerrors.InternalServerException()
	}

	return nil
}

func (ws *WalletService) UpdateBalance(userId int, balance float64) *applicationerrors.ApplicationError {

	repo := ws.getRepository()

	amount := int(balance * 1000)

	err := repo.Update(userId, amount)

	if err != nil {
		return applicationerrors.InternalServerException()
	}

	return nil
}

func (ws *WalletService) GetBalance(userId int) (*models.WalletModel, error) {

	repo := ws.getRepository()

	balance, err := repo.GetByUserId(userId)

	if err != nil {
		return nil, err
	}

	return balance, nil

}

func (ws *WalletService) Delete(userId int) *applicationerrors.ApplicationError {
	repo := ws.getRepository()

	err := repo.Delete(userId)

	if err != nil {
		return applicationerrors.InternalServerException()
	}

	return nil
}

func (ws *WalletService) getRepository() repositories.WalletRepositoryInterface {

	repo, err := ws.uow.GetRepository(context.Background(), "WalletRepository")

	if err != nil {
		return nil
	}

	return repo.(repositories.WalletRepositoryInterface)
}
