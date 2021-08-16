package storage

import (
	"proj1/models"
)

type UserStorage interface {
	CreateUser(models.User) error
	CheckUserExists(string) (models.User, error)
	CheckPassword(string, string) error
}
