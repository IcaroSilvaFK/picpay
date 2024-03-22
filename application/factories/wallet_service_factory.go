package factories

import (
	"context"
	"database/sql"

	"github.com/IcaroSilvaFK/picpay/application/services"
	"github.com/IcaroSilvaFK/picpay/infra/database"
	"github.com/IcaroSilvaFK/picpay/infra/repositories"
	"github.com/IcaroSilvaFK/picpay/pkg/uow"
)

func NewWalletServiceFactory() services.WalletServiceInterface {

	db := database.NewDbConnection(NewDbConfig())

	u := uow.NewUow(context.Background(), db)

	u.Register("WalletRepository", func(tx *sql.Tx) interface{} {

		walletRepository := repositories.NewWalletRepository(db)

		return walletRepository
	})

	walletService := services.NewWalletService(u)

	return walletService
}
