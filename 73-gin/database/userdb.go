package database

import (
	"demo/models"
	"demo/share"
	"errors"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) share.IUser {
	return &UserDB{db: db}
}

func (u *UserDB) Create(user *models.User) (*models.User, error) {
	if u.db == nil {
		return nil, errors.New("nil db object")
	}
	if err := u.db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	tx := u.db.Create(user)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) GetByID(id string) (*models.User, error) {
	user := new(models.User)

	tx := u.db.First(user, id)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
func (u *UserDB) DeleteByID(id string) (int64, error) {

	tx := u.db.Delete(&models.User{}, id)

	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
