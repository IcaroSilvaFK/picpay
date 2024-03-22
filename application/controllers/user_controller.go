package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IcaroSilvaFK/picpay/application/services"
)

type userInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Identifier int    `json:"identifier"`
	Type       int    `json:"type"`
}

type UserController struct {
	usvc services.UserServiceInterface
}

func NewUserController(usvc services.UserServiceInterface) *UserController {
	return &UserController{
		usvc: usvc,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	bd := r.Body

	defer bd.Close()

	var input userInput

	if err := json.NewDecoder(bd).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := uc.usvc.Create(input.Name, input.Email, input.Password, input.Identifier, input.Type)

	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
