package utils

import "golang.org/x/crypto/bcrypt"

func MakeRash(str string) (string, error) {

	s, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(s), nil
}
