package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateExerciseAnswer() []*ds.UserSubmittedAnswerExercises {
	return []*ds.UserSubmittedAnswerExercises{
		{
			ID:                           "f05ca788-374c-463a-9522-02c0b201e9bf",
			CreatedAt:                    time.Now(),
			UpdatedAt:                    time.Now(),
			UserTestSubmissionExerciseID: "b363a709-3425-420e-a57d-3b31f99c8db6",
			QuestionExerciseID:           "47669ec3-0483-4953-bd58-12b4166b9468",
			ChoiceExerciseID:             "90f91f8b-0a1d-4977-a7e8-09d27313cac3",
		},
		{
			ID:                           "443d9b71-b8bb-4cab-9f48-3a58bd04ccdc",
			CreatedAt:                    time.Now(),
			UpdatedAt:                    time.Now(),
			UserTestSubmissionExerciseID: "b363a709-3425-420e-a57d-3b31f99c8db6",
			QuestionExerciseID:           "e1d0f148-1863-4471-b8ab-90982fa70d97",
			ChoiceExerciseID:             "1c4cd26d-bb4f-4f91-963b-1e397a879201",
		},
	}
}
