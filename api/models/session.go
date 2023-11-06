package models

import (
	"encoding/json"
	"errors"
)

type Session struct {
	Username   string `json:"username"`
	Session_ID string `json:"session_id"`
}

func (session *Session) Is_Populated() error {
	if session.Username == "" {
		return errors.New("username is empty")
	}
	if session.Session_ID == "" {
		return errors.New("session_id is empty")
	}
	return nil // All fields are populated
}

// User struct to map
func (session *Session) To_Map() (map[string]any, error) {
	json_data, err := json.Marshal(session)
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
