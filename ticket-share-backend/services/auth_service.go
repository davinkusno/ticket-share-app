package services

import (
	"log"
	"ticket-share-backend/models"

	"gorm.io/gorm"
)

// RegisterUser stores the user with a plain text password in the database
func RegisterUser(db *gorm.DB, user *models.User) error {
	// Directly save the plain text password
	return db.Create(user).Error
}

// CheckPassword compares a plain password with the stored plain text password
func CheckPassword(plainPassword, storedPassword string) bool {
	log.Println(plainPassword)
    log.Println(storedPassword)
	return plainPassword == storedPassword
}

// FindUserByEmail finds a user by their email in the database
func FindUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("Error finding user by email:", err)
		return nil, err
	}
	return &user, nil
}
