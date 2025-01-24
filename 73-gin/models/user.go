package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	//*gorm.Model
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	ContactNo   string `json:"contact_no"`
	Status      string `json:"status"`
	LastUpdated int64  `json:"last_updated"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid name")
	}
	if u.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}

func (u *User) ToBytes() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) ToString() string {
	b, err := u.ToBytes()
	if err != nil {
		return ""
	}
	return string(b)
}
