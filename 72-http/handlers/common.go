package handlers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.Write([]byte("Pong"))
	w.WriteHeader(200)
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.Write([]byte("Ok"))
	w.WriteHeader(200)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("invalid data"))
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := new(models.User)

	err = json.Unmarshal(bytes, user)
	if err != nil {
		w.Write([]byte("invalid data"))
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		w.Write([]byte("invalid data"))
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Status = "active"
	user.LastUpdated = time.Now().Unix()

	//err = utils.SaveUser("users.txt", user)
	utils.ChUser <- user

	// if err != nil {
	// 	w.Write([]byte("invalid data"))
	// 	log.Println(err.Error())
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	w.Write([]byte("User successfully stored "))
	w.WriteHeader(201)
}
