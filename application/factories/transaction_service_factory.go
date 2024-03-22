package factories

import (
	"context"
	"database/sql"

	"github.com/IcaroSilvaFK/picpay/application/services"
	"github.com/IcaroSilvaFK/picpay/infra/configs"
	"github.com/IcaroSilvaFK/picpay/infra/database"
	"github.com/IcaroSilvaFK/picpay/infra/repositories"
	httpclient "github.com/IcaroSilvaFK/picpay/pkg/http"
	"github.com/IcaroSilvaFK/picpay/pkg/uow"
)

func NewTransactionFactory() services.TransactionsServiceInterface {

	db := database.NewDbConnection(NewDbConfig())

	u := uow.NewUow(context.Background(), db)

	u.Register("WalletRepository", func(tx *sql.Tx) interface{} {
		walletRepository := repositories.NewWalletRepository(tx)
		return walletRepository
	})
	u.Register("UserRepository", func(tx *sql.Tx) interface{} {
		repo := repositories.NewUserRepository(tx)
		return repo
	})
	u.Register("TransactionsRepository", func(tx *sql.Tx) interface{} {
		walletRepository := repositories.NewTransactionRepository(tx)
		return walletRepository
	})

	amqp := configs.GetRabbitMQChannel()
	client := httpclient.NewHttpClient("https://run.mocky.io/v3")

	walletService := services.NewTransactionService(u, amqp, client)

	return walletService
}
