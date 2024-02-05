package dto

type UserSubmittedAnswerExercisesRequest struct {
	UserTestSubmissionExerciseID string             `json:"UserTestSubmissionExerciseID" form:"UserTestSubmissionExerciseID"`
	PairOfUserAnswer             []PairOfUserAnswer `json:"PairOfUserAnswer"`
}

type UserSubmittedAnswerExercisesResponse struct {
	Id                           string `json:"Id"`
	UserTestSubmissionExerciseID string `json:"UserTestSubmissionExerciseID" form:"UserTestSubmissionExerciseID"`
	QuestionExerciseID           string `json:"QuestionExerciseID" form:"QuestionExerciseID"`
	ChoiceExerciseID             string `json:"ChoiceExerciseID" form:"ChoiceExerciseID"`
}

type PairOfUserAnswer struct {
	QuestionExerciseID string `json:"QuestionExerciseID" form:"QuestionExerciseID"`
	ChoiceExerciseID   string `json:"ChoiceExerciseID" form:"ChoiceExerciseID"`
}
