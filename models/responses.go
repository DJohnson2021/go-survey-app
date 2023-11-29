package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/DJohnson2021/go-survey-app/db"
	"gorm.io/gorm"
)

type Response struct {
	ID            int32     `gorm:"primaryKey;autoIncrement" db:"id" json:"id"`
	Question_id   int32     `db:"question_id" json:"question_id"`
	User_id       int32     `db:"user_id" json:"user_id"`
	Response string    `db:"response" json:"response"`
	Created_At    time.Time `db:"timestamp" json:"timestamp"`
}

func CreateResponse(response *Response) error {
	// Create a new record in the database
	if err := db.DB.Create(response).Error; err != nil {
		return fmt.Errorf("error creating response: %v", err)
	}
	return nil
}

func DeleteResponseByID(id int32) error {
	result := db.DB.Delete(&Response{}, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Response record with id %v not found", id)
		}
		return result.Error
	}
	return nil
}

func ModifyResponse(response *Response) error {
	result := db.DB.Save(response)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Response record with id %v not found", response.ID)
		}
		return result.Error
	}
	return nil
}