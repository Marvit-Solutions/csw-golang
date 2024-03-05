package dto

type SubjectTestTypeExercisesRequest struct {
	TestType          string                     `json:"TestType" form:"TestType"`
	QuestionExercises []QuestionExercisesRequest `json:"QuestionExercises""`
}

type SubjectTestTypeExercisesResponse struct {
	Id                string                     `json:"id" `
	TestType          string                     `json:"TestType" form:"TestType"`
	QuestionExercises []QuestionExercisesRequest `json:"QuestionExercises,omitempty""`
}
