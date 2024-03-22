package models

import "github.com/IcaroSilvaFK/picpay/application/utils"

type UserModel struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Identifier int
	Type       int
}

const (
	COMMON     = 0
	SHOPKEEPER = 1
)

func NewUser(name, email, password string, identifier, tp int) *UserModel {

	u := &UserModel{
		Name:       name,
		Email:      email,
		Password:   password,
		Identifier: identifier,
		Type:       tp,
	}

	u.makeRash()

	return u
}

func (u *UserModel) makeRash() string {

	u.Password, _ = utils.MakeRash(u.Password)

	return ""
}
