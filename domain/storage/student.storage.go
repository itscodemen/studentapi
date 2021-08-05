package storage

import (
	"proj1/domain/filters"
	"proj1/domain/pagination"
	"proj1/models"
)

type StudentStorage interface {
	InsertStudent(models.Student) (models.Student, error)
	GetStudents(*pagination.Pagination, *filters.StudentFilter) ([]models.Student, error)
	GetStudentByID(string) (models.Student, error)
	UpdateStudent(models.Student, string) (models.Student, error)
	DeleteStudent(models.Student, string) (models.Student, error)
	CheckEmailExists(string) (models.Student, error)
	CheckPhoneExists(string) (models.Student, error)
}
