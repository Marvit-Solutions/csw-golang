package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateGradeQuizzes() []*ds.GradeQuizzes {
	gradeQuizzes := []*ds.GradeQuizzes{
		// User 1 Pretest Pancasila
		{
			ID:                       "58c94d7b-7a0f-44fb-8eca-d67c8a459e9f",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5be14",
			UserID:                   "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			TestTypeQuizID:           "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			Score:                    20,
			GradingTime:              time.Now(),
		},
		// User 2 Pretest Pancasila
		{
			ID:                       "ed0155e5-222c-4014-822d-b6dfd0d99508",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			UserID:                   "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			TestTypeQuizID:           "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			Score:                    20,
			GradingTime:              time.Now(),
		},
	}

	return gradeQuizzes
}
