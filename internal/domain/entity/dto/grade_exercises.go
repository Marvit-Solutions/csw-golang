package dto

import (
	"time"
)

type GradeExercisesRequest struct {
	UserTestSubmissionExerciseID string `json:"UserTestSubmissionExerciseID" form:"UserTestSubmissionExerciseID"`
	UserID                       string `json:"UserID" form:"UserID"`
	TestTypeExerciseID           string `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
}

type GradeExercisesResponse struct {
	ID                           string    `gorm:"type:text;primaryKey"`
	UserTestSubmissionExerciseID string    `json:"UserTestSubmissionExerciseID" form:"UserTestSubmissionExerciseID"`
	UserID                       string    `json:"UserID" form:"UserID"`
	TestTypeExerciseID           string    `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
	Score                        int       `json:"Score" form:"Score"`
	GradingTime                  time.Time `json:"GradingTime" form:"GradingTime"`
}
