package repository

import (
	"github.com/felipelaptrin/go-api/pkg/config"
	"github.com/felipelaptrin/go-api/pkg/models"
)

func CheckUserExists(id int) bool {
	var user models.User
	err := config.Db.First(&user, id).Error

	return err == nil
}

func GetUserById(id int) (user *models.User, err error) {
	var _user models.User
	_err := config.Db.First(&_user, id)

	return &_user, _err.Error
}

func GetAllUsers() (user *[]models.User, err error) {
	var _users []models.User
	_err := config.Db.Find(&_users)

	return &_users, _err.Error
}

func CreateUser(user *models.User) error {
	err := config.Db.Create(&user)

	return err.Error
}

func UpdateUserById(id int, user *models.User) error {
	err := config.Db.Where("id = ?", id).Updates(&user)

	return err.Error
}

func DeleteUserById(id int) (user *models.User, err error) {
	var _user models.User

	_err := config.Db.Where("id = ?", id).Delete(&_user).Error
	return &_user, _err
}
