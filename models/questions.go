package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/DJohnson2021/go-survey-app/db"
	"gorm.io/gorm"
)

type Question struct {
	ID            int32     `gorm:"primaryKey" gorm:"autoIncrement" db:"id" json:"id"`
	Question_text string    `db:"question_text" json:"question_text"`
	Survey_id     int32     `db:"survey_id" json:"survey_id"`
	Created_At    time.Time `db:"timestamp" json:"timestamp"`
}

func CreateQuestion(question *Question) error {
	// Create a new record in the database
	if err := db.DB.Create(question).Error; err != nil {
		return fmt.Errorf("error creating question: %v", err)
	}
	return nil
}

func DeleteQuestionByID(id int32) error {
	result := db.DB.Delete(&Question{}, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Question record with id %v not found", id)
		}
		return result.Error
	}
	return nil
}

func ModifyQuestion(question *Question) error {
	result := db.DB.Save(question)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Question record with id %v not found", question.ID)
		}
		return result.Error
	}
	return nil
}