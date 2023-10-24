package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/DJohnson2021/go-survey-app/db"
	"gorm.io/gorm"
)

type User struct {
	ID         int32     `db:"id" json:"id"`
	GoogleID   string    `db:"google_id" json:"google_id"`
	Username   string    `db:"username" json:"username"`
	GivenName  string    `db:"given_name" json:"given_name"`
	FamilyName string    `db:"family_name" json:"family_name"`
	Email      string    `db:"email" json:"email"`
	Created_At time.Time `db:"timestamp" json:"timestamp"`
	// ... other fields
}

func CreateUser(user *User) error {
	// Create a new record in the database
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(google_id string) (*User, error) {
	var user User
	result := db.DB.First(&user, "google_id = ?", google_id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User record with google_id %v not found", google_id)
		}
		return nil, result.Error
	}
	return &user, nil
}

func DeleteUserByID(google_id string) error {
	result := db.DB.Delete(&User{}, "google_id = ?", google_id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("User record with google_id %v not found", google_id)
		}
		return result.Error
	}
	return nil
}
