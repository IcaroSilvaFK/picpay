package factories

import (
	"context"
	"database/sql"

	"github.com/IcaroSilvaFK/picpay/application/services"
	"github.com/IcaroSilvaFK/picpay/infra/database"
	"github.com/IcaroSilvaFK/picpay/infra/repositories"
	"github.com/IcaroSilvaFK/picpay/pkg/uow"
)

func NewUserServiceFactory() services.UserServiceInterface {
	db := database.NewDbConnection(NewDbConfig())

	u := uow.NewUow(context.Background(), db)

	u.Register("UserRepository", func(tx *sql.Tx) interface{} {
		repo := repositories.NewUserRepository(tx)
		return repo
	})

	return services.NewUserService(u)
}
