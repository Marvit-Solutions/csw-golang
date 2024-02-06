package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateExerciseSubmission() []*ds.UserTestSubmissionExercises {
	return []*ds.UserTestSubmissionExercises{{
		ID:                 "b363a709-3425-420e-a57d-3b31f99c8db6",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		UserID:             "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		TestTypeExerciseID: "9975116b-01a4-4978-90aa-a3a7525d504a",
		SubmissionTIme:     time.Now(),
	}}
}
