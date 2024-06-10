package repository

import (
	"HSAuth/models"
	"HSAuth/store"
	"gorm.io/gorm"
)

func DoesEmailExist(email string) (bool, error) { //check if email already exists
	var user models.User
	err := store.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
