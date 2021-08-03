package repository

import (
	"errors"
	"fmt"
	"proj1/models"
	"strconv"

	"gorm.io/gorm"
)

type StudentPersistStorage struct {
	db *gorm.DB
}

func (s *StudentPersistStorage) InsertStudent(st models.Student) (models.Student, error) {
	err := s.db.Create(&st).Error
	if err != nil {
		return models.Student{}, nil
	}
	return st, nil
}

func (s *StudentPersistStorage) GetStudent() ([]models.Student, error) {
	var Student []models.Student
	err := s.db.Find(&Student).Error
	if err != nil {
		return []models.Student{}, nil
	}
	return Student, nil
}

func (s *StudentPersistStorage) GetStudentByID(id string) (models.Student, error) {
	var st models.Student
	err := s.db.Where("id = ?", id).First(&st).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Student{}, err
		}
		return models.Student{}, err
	}
	return st, nil
}

func (s *StudentPersistStorage) UpdateStudent(st models.Student, id string) (models.Student, error) {
	fmt.Println(st)
	var student models.Student
	err := s.db.Where("id = ?", id).First(&student).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Student{}, err
		}
		return models.Student{}, err
	}
	st.ID, _ = strconv.Atoi(id)
	s.db.Save(&st)
	return st, nil

}

func (s *StudentPersistStorage) DeleteStudent(st models.Student, id string) (models.Student, error) {
	err := s.db.Where("id = ?", id).Delete(&st).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Student{}, err
		}
		return models.Student{}, err
	}
	return st, nil
}

func (s *StudentPersistStorage) CheckEmailExists(email string) (models.Student, error) {
	var st models.Student
	err := s.db.Where("email = ?", email).First(&st).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Student{}, err
		}
		return models.Student{}, err
	}
	return st, nil
}

func (s *StudentPersistStorage) CheckPhoneExists(phone string) (models.Student, error) {
	var st models.Student
	err := s.db.Where("phone = ?", phone).First(&st).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Student{}, err
		}
		return models.Student{}, err
	}
	return st, nil
}

func NewStudentStorage(db *gorm.DB) *StudentPersistStorage {
	s := &StudentPersistStorage{db: db}
	return s
}
