package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IcaroSilvaFK/picpay/application/services"
)

type walletInput struct {
	UserId  int     `json:"user_id,omitempty"`
	Balance float64 `json:"balance"`
}

type walletOutput struct {
	UserId  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}

type WalletController struct {
	svc services.WalletServiceInterface
}

type WalletControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByUser(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewWalletController(svc services.WalletServiceInterface) WalletControllerInterface {
	return &WalletController{
		svc: svc,
	}
}

func (wc *WalletController) Create(w http.ResponseWriter, r *http.Request) {

	bd := r.Body

	defer bd.Close()

	var input walletInput

	if err := json.NewDecoder(bd).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := wc.svc.Create(input.UserId, input.Balance)

	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (wc *WalletController) GetByUser(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Query().Get("userId")

	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, "user_id must be a number", http.StatusBadRequest)
		return
	}

	wallet, err := wc.svc.GetBalance(id)

	if err != nil {
		return
	}

	out := walletOutput{
		UserId:  wallet.UserId,
		Balance: float64(wallet.Amount),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(out)

}

func (wc *WalletController) Update(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Query().Get("userId")

	id, err := strconv.Atoi(userId)

	if err != nil {

		http.Error(w, "userId must be a number", http.StatusBadRequest)
		return
	}

	var input walletInput

	bd := r.Body

	defer bd.Close()

	if err := json.NewDecoder(bd).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	appErr := wc.svc.UpdateBalance(id, input.Balance)

	if appErr != nil {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (wc *WalletController) Delete(w http.ResponseWriter, r *http.Request) {
	walletId := r.URL.Query().Get("walletId")

	id, err := strconv.Atoi(walletId)

	if err != nil {
		http.Error(w, "walletId must be a number", http.StatusBadRequest)
		return
	}

	appErr := wc.svc.Delete(id)

	if appErr != nil {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
