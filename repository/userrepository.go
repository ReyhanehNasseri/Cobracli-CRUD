package repository

import (
	"learn_1/model"
)

type UserRepository interface {
	Connect() error
	InsertUser(user model.User) (model.User, error)
	Read_all()
	Read(id string) error
	UpdateUser(id string, newfirstname string, newlastname string) error
	DeleteUser(id string) error
}
