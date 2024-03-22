package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IcaroSilvaFK/picpay/application/services"
)

type TransactionController struct {
	svc services.TransactionsServiceInterface
}

type transactionInput struct {
	Value float64 `json:"value"`
	Payer int     `json:"payer"`
	Payee int     `json:"payee"`
}

type TransactionControllerInterface interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewTransactionController(
	svc services.TransactionsServiceInterface,
) TransactionControllerInterface {

	return &TransactionController{
		svc,
	}
}

func (tc *TransactionController) Execute(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	bd := r.Body

	defer bd.Close()

	bt, err := io.ReadAll(bd)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": %v}`, err)))
		return
	}

	var input transactionInput

	if err := json.Unmarshal(bt, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": %v}`, err)))
		return
	}

	appErr := tc.svc.Create(input.Payer, input.Payee, input.Value)

	if appErr != nil {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
