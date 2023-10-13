package models

import (
	"errors"
	"time"

	"github.com/DJohnson2021/go-survey-app/db"
	"gorm.io/gorm"
)

type User struct {
    ID        int32     `db:"id" json:"id"`
    GoogleID  int64     `db:"google_id" json:"google_id"`
    Username  string    `db:"username" json:"username"`
    Email     string    `db:"email" json:"email"`
    Timestamp time.Time `db:"timestamp" json:"timestamp"`
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
            return nil, nil
        }
        return nil, result.Error
    }
    return &user, nil
}

