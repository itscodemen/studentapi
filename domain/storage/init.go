package storage

import (
	"proj1/domain/storage/repository"

	"gorm.io/gorm"
)

var Student StudentStorage

func InitStorage(db *gorm.DB) error {
	Student = repository.NewStudentStorage(db)
	return nil
}
