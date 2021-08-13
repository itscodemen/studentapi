package storage

import (
	"proj1/domain/storage/repository"

	"gorm.io/gorm"
)

var (
	Student StudentStorage
	User    UserStorage
)

func InitStorage(db *gorm.DB) error {
	Student = repository.NewStudentStorage(db)
	User = repository.NewUserStorage(db)
	return nil
}
