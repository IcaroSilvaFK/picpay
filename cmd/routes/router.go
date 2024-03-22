package routes

import (
	"net/http"

	"github.com/IcaroSilvaFK/picpay/application/controllers"
	"github.com/IcaroSilvaFK/picpay/application/factories"
)

func NewApiRouter(mx *http.ServeMux) {

	uController := controllers.NewUserController(factories.NewUserServiceFactory())
	wController := controllers.NewWalletController(factories.NewWalletServiceFactory())
	transctionController := controllers.NewTransactionController(factories.NewTransactionFactory())

	mx.HandleFunc("POST /users", uController.CreateUser)
	mx.HandleFunc("POST /wallets", wController.Create)
	mx.HandleFunc("POST /transactions", transctionController.Execute)
}
