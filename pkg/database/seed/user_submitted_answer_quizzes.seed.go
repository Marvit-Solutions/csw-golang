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
			ID:                       "aa6d085b-71ee-416d-8777-4de3a26d01b1",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "69551cd4-47f1-4fc1-b00f-a7df35a90077",
			ChoiceQuizID:             "25916cff-5cee-4cd3-b90b-0893c899dbc6",
		},
		//  Nomor 2
		{
			ID:                       "0d82191e-a0ee-4d66-974c-3abfc0f2a14e",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "53ed22ce-0af5-453b-b87c-0bfb1565556a",
			ChoiceQuizID:             "224892d0-a654-4f2a-ac63-3b81ca6939fd",
		},
		//  Nomor 3
		{
			ID:                       "e13322f4-d292-4e99-88e4-2f9b785578b8",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "4f7c1f79-25fd-4e60-9219-073a414e9720",
			ChoiceQuizID:             "e2c70691-d535-4f43-b6b0-8d461a8b48c4",
		},
		//  Nomor 4
		{
			ID:                       "25e5649f-6fb1-4ae3-ae8c-68cbb31e5b93",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "0d1f4266-ad1e-41f0-b36a-9dc808e63c93",
			ChoiceQuizID:             "5bdac59c-1f7d-4551-9534-c8e877277dc2",
		},
		//  Nomor 5
		{
			ID:                       "e381d26e-77b2-46ba-a253-1df16bf5c1d1",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540bea54",
			QuestionQuizID:           "0c301ace-1869-475c-99a4-7585a509e622",
			ChoiceQuizID:             "25838b2c-00a7-4abd-ace7-585f00ca2a7c",
		},
		// MODULE TEST PANCASILA USER 1
		//  Nomor 1 (benar)
		{
			ID:                       "5cb8f7af-d36c-4906-97e3-bc84691d1123",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5b123",
			QuestionQuizID:           "2b45009a-e8cc-43f9-b957-c68fe137f9c3",
			ChoiceQuizID:             "94935ab4-8b78-4b93-a26f-12f161a58107",
		},
		//  Nomor 2 (salah)
		{
			ID:                       "41a9b232-cb23-4661-b48a-44ce02d37124",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5b123",
			QuestionQuizID:           "f4fad3b8-e6f0-48b3-8dd9-ba6982b90634",
			ChoiceQuizID:             "f028ebbe-3d19-41db-9df5-4893ddbff621",
		},
		//  Nomor 3 (benar)
		{
			ID:                       "f07eaf84-d0eb-4da3-ba39-59abe24f0125",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5b123",
			QuestionQuizID:           "3fdcf4ac-e8ba-4bc8-bbdf-906fc756ebc2",
			ChoiceQuizID:             "e2c70691-d535-4f43-b6b0-8d461a8b48c1",
		},
		//  Nomor 4 (salah)
		{
			ID:                       "dfe9c8a6-8c1e-4c7e-82c7-f0928a975126",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5b123",
			QuestionQuizID:           "b8456d7e-6940-485b-b32c-16207d8e1531",
			ChoiceQuizID:             "43d60a13-01c4-4988-833f-d53483b6a112",
		},
		//  Nomor 5 (benar)
		{
			ID:                       "3215de91-5055-4b37-81b2-b872ec078127",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "3611a539-5252-4f64-b70f-082dd0d5b123",
			QuestionQuizID:           "fe2d281e-fc4e-427e-b487-63d4065ef7a3",
			ChoiceQuizID:             "9cedbb63-5e07-4b35-833b-fe1337e2d833",
		},
		// MODULE TEST PANCASILA USER 1
		//  Nomor 1 (benar)
		{
			ID:                       "5cb8f7af-d36c-4906-97e3-bc84691d1128",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540be124",
			QuestionQuizID:           "2b45009a-e8cc-43f9-b957-c68fe137f9c3",
			ChoiceQuizID:             "94935ab4-8b78-4b93-a26f-12f161a58107",
		},
		//  Nomor 2 (benar)
		{
			ID:                       "41a9b232-cb23-4661-b48a-44ce02d37129",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540be124",
			QuestionQuizID:           "f4fad3b8-e6f0-48b3-8dd9-ba6982b90634",
			ChoiceQuizID:             "224892d0-a654-4f2a-ac63-3b81ca6939fd",
		},
		//  Nomor 3 (benar)
		{
			ID:                       "f07eaf84-d0eb-4da3-ba39-59abe24f0130",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540be124",
			QuestionQuizID:           "3fdcf4ac-e8ba-4bc8-bbdf-906fc756ebc2",
			ChoiceQuizID:             "e2c70691-d535-4f43-b6b0-8d461a8b48c4",
		},
		//  Nomor 4 (salah)
		{
			ID:                       "dfe9c8a6-8c1e-4c7e-82c7-f0928a975131",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540be124",
			QuestionQuizID:           "b8456d7e-6940-485b-b32c-16207d8e1531",
			ChoiceQuizID:             "43d60a13-01c4-4988-833f-d53483b6a112",
		},
		//  Nomor 5 (benar)
		{
			ID:                       "3215de91-5055-4b37-81b2-b872ec078132",
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
			UserTestSubmissionQuizID: "58bbf5b2-1838-4a68-83b7-0d8c540be124",
			QuestionQuizID:           "fe2d281e-fc4e-427e-b487-63d4065ef7a3",
			ChoiceQuizID:             "9cedbb63-5e07-4b35-833b-fe1337e2d833",
		},
	}

	return userSubmittedAnswerQuizzes
}
