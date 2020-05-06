package controllers

import (
	"fmt"
	"fyture/models"

	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileImage"`
}

func CreateUser(input *CreateUserInput, db *gorm.DB) models.User {
	// Create Feature
	user := models.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		ProfileImage: input.ProfileImage,
	}
	db.Create(&user)

	return user
}

func FindUserByID(id uint, db *gorm.DB) (models.User, error) {
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, fmt.Errorf("User not found: %s", err.Error())
	}

	return user, nil
}

func FindUserByEmail(email string, db *gorm.DB) (models.User, error) {
	// Get model if exist
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, fmt.Errorf("User not found: %s", err.Error())
	}

	return user, nil
}
