package services

import (
	"context"
	"database/sql"
	"fmt"

	applicationerrors "github.com/IcaroSilvaFK/picpay/application/errors"
	"github.com/IcaroSilvaFK/picpay/infra/models"
	"github.com/IcaroSilvaFK/picpay/infra/repositories"
	"github.com/IcaroSilvaFK/picpay/pkg/uow"
)

type UserService struct {
	uow uow.UowInterface
}

type UserServiceInterface interface {
	Create(name, email, password string, identifier, tp int) *applicationerrors.ApplicationError
	FindById(id int) (*models.UserModel, error)
	Delete(id int) *applicationerrors.ApplicationError
}

func NewUserService(
	uow uow.UowInterface,
) UserServiceInterface {

	return &UserService{
		uow: uow,
	}
}

func (us *UserService) Create(name, email, password string, identifier, tp int) *applicationerrors.ApplicationError {

	userRepo := us.getUserRepo()

	u := models.NewUser(name, email, password, identifier, tp)

	err := userRepo.Create(u)
	if err != nil {
		return applicationerrors.InternalServerException()
	}
	return nil

}

func (us *UserService) FindById(id int) (*models.UserModel, error) {

	uRepo := us.getUserRepo()

	return uRepo.GetById(id)
}

func (us *UserService) Delete(id int) *applicationerrors.ApplicationError {

	uRepo := us.getUserRepo()

	err := uRepo.Delete(id)

	if err == sql.ErrNoRows {
		return applicationerrors.NotFoundException(fmt.Sprintf("Record id %d not found", id))
	}

	if err != nil {
		return applicationerrors.InternalServerException()
	}
	return nil

}

func (us *UserService) getUserRepo() repositories.UserRepositoryInterface {

	repo, err := us.uow.GetRepository(context.Background(), "UserRepository")

	if err != nil {
		panic(err)
	}

	return repo.(repositories.UserRepositoryInterface)

}
