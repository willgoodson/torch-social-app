package utility

import (
	"errors"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

func Hash_Password(password string) (string, error) {
	password_bytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(password_bytes, 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Check_Password(test_pass string, stored_pass string) error {
	// Ensure Store Password Is A String
	if reflect.TypeOf(stored_pass).Kind() != reflect.String {
		return errors.New("stored_pass is not string")
	}
	err := bcrypt.CompareHashAndPassword([]byte(stored_pass), []byte(test_pass))
	if err != nil {
		return errors.New("password does not match")
	}
	return nil
}
