package repository

import (
	"proj1/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *StudentPersistStorage) CreateUser(user models.User) error {
	err := s.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentPersistStorage) CheckUserExists(payload string) (models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", payload).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *StudentPersistStorage) CheckPassword(providedPassword string, providedEmail string) error {
	var user models.User
	var res string
	s.db.Model(&user).Where("email = ?", providedEmail).Pluck("password", &res)
	err := bcrypt.CompareHashAndPassword([]byte(res), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil

}

func NewUserStorage(db *gorm.DB) *StudentPersistStorage {
	s := &StudentPersistStorage{db: db}
	return s
}
