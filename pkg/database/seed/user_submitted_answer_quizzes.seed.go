package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateUserSubmittedAnswerQuizzes() []*ds.UserSubmittedAnswerQuizzes {
	userSubmittedAnswerQuizzes := []*ds.UserSubmittedAnswerQuizzes{
		// PRETEST PANCASILA USER 1
		//  Nomor 1
		{
			ID:                       "5cb8f7af-d36c-4906-97e3-bc84691d1e19",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5be14",
			QuestionQuizID:           "69551cd4-47f1-4fc1-b00f-a7df35a90077",
			ChoiceQuizID:             "25916cff-5cee-4cd3-b90b-0893c899dbc6",
		},
		//  Nomor 2
		{
			ID:                       "41a9b232-cb23-4661-b48a-44ce02d370e9",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5be14",
			QuestionQuizID:           "53ed22ce-0af5-453b-b87c-0bfb1565556a",
			ChoiceQuizID:             "224892d0-a654-4f2a-ac63-3b81ca6939fd",
		},
		//  Nomor 3
		{
			ID:                       "f07eaf84-d0eb-4da3-ba39-59abe24f0a97",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5be14",
			QuestionQuizID:           "4f7c1f79-25fd-4e60-9219-073a414e9720",
			ChoiceQuizID:             "e2c70691-d535-4f43-b6b0-8d461a8b48c4",
		},
		//  Nomor 4
		{
			ID:                       "dfe9c8a6-8c1e-4c7e-82c7-f0928a975fec",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5be14",
			QuestionQuizID:           "0d1f4266-ad1e-41f0-b36a-9dc808e63c93",
			ChoiceQuizID:             "5bdac59c-1f7d-4551-9534-c8e877277dc2",
		},
		//  Nomor 5
		{
			ID:                       "3215de91-5055-4b37-81b2-b872ec078ac9",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5be14",
			QuestionQuizID:           "0c301ace-1869-475c-99a4-7585a509e622",
			ChoiceQuizID:             "25838b2c-00a7-4abd-ace7-585f00ca2a7c",
		},
		// PRETEST PANCASILA USER 2
		//  Nomor 1
		{
			ID:                       "60beba78-54aa-4383-b1c2-361baa360f7a",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "69551cd4-47f1-4fc1-b00f-a7df35a90077",
			ChoiceQuizID:             "25916cff-5cee-4cd3-b90b-0893c899dbc6",
		},
		//  Nomor 2
		{
			ID:                       "41a9b232-cb23-4661-b48a-44ce02d370e9",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "53ed22ce-0af5-453b-b87c-0bfb1565556a",
			ChoiceQuizID:             "224892d0-a654-4f2a-ac63-3b81ca6939fd",
		},
		//  Nomor 3
		{
			ID:                       "f07eaf84-d0eb-4da3-ba39-59abe24f0a97",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "4f7c1f79-25fd-4e60-9219-073a414e9720",
			ChoiceQuizID:             "e2c70691-d535-4f43-b6b0-8d461a8b48c4",
		},
		//  Nomor 4
		{
			ID:                       "dfe9c8a6-8c1e-4c7e-82c7-f0928a975fec",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "0d1f4266-ad1e-41f0-b36a-9dc808e63c93",
			ChoiceQuizID:             "5bdac59c-1f7d-4551-9534-c8e877277dc2",
		},
		//  Nomor 5
		{
			ID:                       "3215de91-5055-4b37-81b2-b872ec078ac9",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "0c301ace-1869-475c-99a4-7585a509e622",
			ChoiceQuizID:             "25838b2c-00a7-4abd-ace7-585f00ca2a7c",
		},
	}

	return userSubmittedAnswerQuizzes
}
