package models

import (
	"encoding/json"
	"errors"
)

type UserLogin struct {
	Email     string `json:"email"`
	Passwortd string `json:"password"`
}

func (u *UserLogin) Validate() error {
	if u.Passwortd == "" {
		return errors.New("invalid password")
	}
	if u.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}

func (u *UserLogin) ToBytes() ([]byte, error) {
	return json.Marshal(u)
}

func (u *UserLogin) ToString() string {
	b, err := u.ToBytes()
	if err != nil {
		return ""
	}
	return string(b)
}
