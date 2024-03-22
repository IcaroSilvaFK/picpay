package services

import (
	"context"
	"encoding/json"

	applicationerrors "github.com/IcaroSilvaFK/picpay/application/errors"
	"github.com/IcaroSilvaFK/picpay/infra/models"
	"github.com/IcaroSilvaFK/picpay/infra/repositories"
	httpclient "github.com/IcaroSilvaFK/picpay/pkg/http"
	"github.com/IcaroSilvaFK/picpay/pkg/uow"
	"github.com/streadway/amqp"
)

type TransactionsService struct {
	uow        uow.UowInterface
	amqp       *amqp.Channel
	httpclient httpclient.HttpClientInterface
}

type TransactionsServiceInterface interface {
	Create(payer, payee int, amount float64) *applicationerrors.ApplicationError
	GetMyTransactions(userId int) ([]*models.TransactionModel, *applicationerrors.ApplicationError)
	Delete(id string) *applicationerrors.ApplicationError
}

func NewTransactionService(
	uow uow.UowInterface,
	amqp *amqp.Channel,
	httpclient httpclient.HttpClientInterface,
) TransactionsServiceInterface {
	return &TransactionsService{
		uow, amqp, httpclient,
	}
}

func (ts *TransactionsService) Create(payer, payee int, amount float64) *applicationerrors.ApplicationError {

	return ts.uow.Do(context.Background(), func(uow *uow.Uow) *applicationerrors.ApplicationError {

		walletrepo := ts.getWalletRepository()
		userrepo := ts.getUserRepository()

		sender, err := userrepo.GetById(payer)

		value := int(amount * 1000)

		if err != nil {
			return applicationerrors.InternalServerException()
		}

		if sender.Type == models.SHOPKEEPER {
			return applicationerrors.BadRequestException("Users type Shopkeeper not available from transaction")
		}

		currentwallet, err := walletrepo.GetByUserId(payer)

		if err != nil {
			return applicationerrors.InternalServerException()
		}

		if currentwallet.Amount < value {
			return applicationerrors.BadRequestException("Your current amount not valid from execute transaction please deposit more.")
		}

		var apiResponse struct {
			Message string `json:"message"`
		}

		if err := ts.httpclient.Get("/5794d450-d2e2-4412-8131-73d0293ac1cc", &apiResponse); err != nil {
			return applicationerrors.InternalServerException()
		}

		if apiResponse.Message != "Autorizado" {
			return applicationerrors.BadRequestException("Your transaction is not authorized")
		}

		currentwallet.Amount -= value

		if err := walletrepo.Update(sender.ID, currentwallet.Amount); err != nil {
			return applicationerrors.InternalServerException()
		}

		revicerwallet, err := walletrepo.GetByUserId(payee)

		if err != nil {
			return applicationerrors.InternalServerException()
		}

		revicerwallet.Amount += value

		if err := walletrepo.Update(payee, revicerwallet.Amount); err != nil {
			return applicationerrors.InternalServerException()
		}

		t := models.NewTransactionModel(payer, payee, value)

		repo := ts.getTransactionRepository()

		err = repo.Create(t)

		if err != nil {
			return applicationerrors.InternalServerException()
		}

		receiver, err := userrepo.GetById(payee)

		if err != nil {
			return applicationerrors.InternalServerException()
		}

		msgPayload := struct {
			PayerEmail string `json:"payer_email"`
			PayeeEmail string `json:"payee_email"`
		}{
			PayerEmail: sender.Email,
			PayeeEmail: receiver.Email,
		}

		jsonOut, _ := json.Marshal(msgPayload)

		msgQueue := amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonOut,
		}

		ts.amqp.Publish(
			"",
			"picpay",
			false,
			false,
			msgQueue,
		)

		return nil
	})
}

func (ts *TransactionsService) GetMyTransactions(userId int) ([]*models.TransactionModel, *applicationerrors.ApplicationError) {

	repo := ts.getTransactionRepository()

	r, err := repo.GetMyTransactions(userId)

	if err != nil {
		return nil, applicationerrors.InternalServerException()
	}

	return r, nil
}

func (ts *TransactionsService) Delete(id string) *applicationerrors.ApplicationError {

	repo := ts.getTransactionRepository()

	err := repo.Delete(id)

	if err != nil {
		return applicationerrors.InternalServerException()
	}

	return nil
}

func (ts *TransactionsService) getTransactionRepository() repositories.TransactionRepositoryInterface {
	repo, err := ts.uow.GetRepository(context.Background(), "TransactionsRepository")

	if err != nil {
		return nil
	}

	return repo.(repositories.TransactionRepositoryInterface)
}
func (ts *TransactionsService) getWalletRepository() repositories.WalletRepositoryInterface {

	repo, err := ts.uow.GetRepository(context.Background(), "WalletRepository")

	if err != nil {
		return nil
	}

	return repo.(repositories.WalletRepositoryInterface)
}

func (ts *TransactionsService) getUserRepository() repositories.UserRepositoryInterface {
	repo, err := ts.uow.GetRepository(context.Background(), "UserRepository")

	if err != nil {
		return nil
	}

	return repo.(repositories.UserRepositoryInterface)
}
