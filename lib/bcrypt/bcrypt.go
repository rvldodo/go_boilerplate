package bcrypt

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/rvldodo/boilerplate/lib/log"
)

func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("Hashing error: %v", err)
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(password string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), plain)
	return err == nil
}
