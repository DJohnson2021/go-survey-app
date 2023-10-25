package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/DJohnson2021/go-survey-app/db"
	"gorm.io/gorm"
)

type Survey struct {
	ID          int32     `gorm:"primaryKey" gorm:"autoIncrement" db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Created_At  time.Time `db:"timestamp" json:"timestamp"`
	Description string    `db:"description" json:"description"`
	User_id     int32     `db:"user_id" json:"user_id"`
}

func CreateSurvey(survey *Survey) error {
	// Create a new record in the database
	if err := db.DB.Create(survey); err != nil {
		return fmt.Errorf("error creating survey: %v", err)
	}
	return nil
}

func DeleteSurveyByID(id int32) error {
	result := db.DB.Delete(&Survey{}, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Survey record with id %v not found", id)
		}
		return result.Error
	}
	return nil
}

func ModifySurvey(survey *Survey) error {
	result := db.DB.Save(survey)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Survey record with id %v not found", survey.ID)
		}
		return result.Error
	}
	return nil
}
