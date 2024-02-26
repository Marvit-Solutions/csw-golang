package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateUserTestSubmissionQuizzes() []*ds.UserTestSubmissionQuizzes {
	userTestSubmissionQuizzes := []*ds.UserTestSubmissionQuizzes{
		// Pretest Pancasila User 1
		{
			ID:             "3611a539-5252-4f64-b70f-082dd0d5be14",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserID:         "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			TestTypeQuizID: "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			SubmissionTIme: time.Now(),
		},
		// Pretest Pancasila User 2
		{
			ID:             "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserID:         "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			TestTypeQuizID: "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			SubmissionTIme: time.Now(),
		},
		//  Module Test Pancasila User 1
		{
			ID:             "3611a539-5252-4f64-b70f-082dd0d5b123",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserID:         "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			TestTypeQuizID: "f8259b82-41ca-4d58-90e7-c2be3f4c235f",
			SubmissionTIme: time.Now(),
		},
		// Module Test Pancasila User 2
		{
			ID:             "58bbf5b2-1838-4a68-83b7-0d8c540be124",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserID:         "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			TestTypeQuizID: "f8259b82-41ca-4d58-90e7-c2be3f4c235f",
			SubmissionTIme: time.Now(),
		},
	}

	return userTestSubmissionQuizzes
}
