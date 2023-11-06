package models

import (
	"encoding/json"
	"errors"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *Login) Is_Populated() error {
	if login.Username == "" {
		return errors.New("username is empty")
	}
	if login.Password == "" {
		return errors.New("password is empty")
	}
	return nil // All fields are populated
}

// User struct to map
func (login *Login) To_Map() (map[string]any, error) {
	json_data, err := json.Marshal(login)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	err = json.Unmarshal(json_data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
