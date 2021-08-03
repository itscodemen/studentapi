package storage

import (
	"proj1/models"
)

type StudentStorage interface {
	InsertStudent(models.Student) (models.Student, error)
	GetStudent(string, string) ([]models.Student, error)
	GetStudentByID(string) (models.Student, error)
	UpdateStudent(models.Student, string) (models.Student, error)
	DeleteStudent(models.Student, string) (models.Student, error)
	CheckEmailExists(string) (models.Student, error)
	CheckPhoneExists(string) (models.Student, error)
}
