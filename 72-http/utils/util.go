package utils

import (
	"demo/models"
	"log"
	"os"
)

var (
	ChUser   chan *models.User
	ChErr    chan error
	FileName string
)

func init() {
	if ChUser == nil {
		ChUser = make(chan *models.User, 10)
	}
	if ChErr == nil {
		ChErr = make(chan error)
	}
	FileName = "users.txt"
	go SaveUserCh()
	go OnError()
}

func SaveUser(filename string, user *models.User) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := user.ToBytes()
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(string(bytes) + "\n"))
	if err != nil {
		return os.ErrClosed
	}
	return nil
}

func SaveUserCh() {
	f, err := os.OpenFile(FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		ChErr <- err
		//return err
	}
	defer f.Close()

	for user := range ChUser {
		bytes, err := user.ToBytes()
		if err != nil {
			ChErr <- err
		}
		_, err = f.Write([]byte(string(bytes) + "\n"))
		if err != nil {
			ChErr <- err
		}
	}
}

func OnError() {
	for err := range ChErr {
		log.Println(err.Error())
	}
}
