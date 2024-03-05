package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubjectTestTypeExercises struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//SubjectID                   string                        `json:"SubjectID" form:"SubjectID"`
	TestType                    string                        `json:"TestType" form:"TestType"`
	GradeExercises              GradeExercises                `gorm:"foreignKey:TestTypeExerciseID"`
	QuestionExercises           []QuestionExercises           `gorm:"foreignKey:TestTypeExerciseID"`
	UserTestSubmissionExercises []UserTestSubmissionExercises `gorm:"foreignKey:TestTypeExerciseID"`
}
