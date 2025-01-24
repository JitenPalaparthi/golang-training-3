package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RETRY_TIMES    = 10
	REGRY_DURATION = time.Second * 5
)

func GetConnection(dsn string) (*gorm.DB, error) {
	count := 0
retry:
	count++
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("trying to connect to the database ---->", count, " times")
		if count <= RETRY_TIMES {
			time.Sleep(REGRY_DURATION)
			goto retry
		}
	}
	return db, err
}

//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
