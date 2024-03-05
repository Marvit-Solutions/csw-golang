package dto

import (
	"time"
)

type UserTestSubmissionExercisesRequest struct {
	UserID             string `json:"UserID" form:"UserID"`
	TestTypeExerciseID string `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
}

type UserTestSubmissionExercisesResponse struct {
	ID                 string    `gorm:"type:text;primaryKey"`
	UserID             string    `json:"UserID" form:"UserID"`
	TestTypeExerciseID string    `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
	SubmissionTime     time.Time `json:"SubmissionTIme" form:"SubmissionTIme"`
}
