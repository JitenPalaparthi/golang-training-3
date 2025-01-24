package share

import "demo/models"

type IUser interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	DeleteByID(id string) (int64, error)
}
