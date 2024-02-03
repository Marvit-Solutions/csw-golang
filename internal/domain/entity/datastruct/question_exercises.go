package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type QuestionExercises struct {
	ID                           string                         `gorm:"type:text;primaryKey"`
	CreatedAt                    time.Time                      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                    time.Time                      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                    gorm.DeletedAt                 `gorm:"index"`
	TestTypeExerciseID           string                         `json:"TestTypeExerciseID" form:"TestTypeExerciseID" gorm:"default:null"`
	Content                      string                         `json:"Content" form:"Content"`
	Weight                       int                            `json:"Weight" form:"Weight"`
	ChoiceExercises              []ChoiceExercises              `gorm:"foreignKey:QuestionExerciseID"`
	UserSubmittedAnswerExercises []UserSubmittedAnswerExercises `gorm:"foreignKey:QuestionExerciseID"`
}
