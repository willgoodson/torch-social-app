package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Is_Populated() error {
	if user.Email == "" {
		return errors.New("email is empty")
	}
	if user.Username == "" {
		return errors.New("username is empty")
	}
	if user.Name == "" {
		return errors.New("name is empty")
	}
	if user.Password == "" {
		return errors.New("password is empty")
	}

	return nil // All fields are populated
}

// User struct to map
func (user *User) To_Map() (map[string]any, error) {
	json_data, err := json.Marshal(user)
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
