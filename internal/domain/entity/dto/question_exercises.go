package dto

type QuestionExercisesRequest struct {
	TestTypeExerciseID string                   `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
	Content            string                   `json:"Content" form:"Content"`
	Weight             int                      `json:"Weight" form:"Weight"`
	Tags               []string                 `json:"Tags" form:"Tags"`
	ChoiceExercises    []ChoiceExercisesRequest `json:"ChoiceExercises"`
}

type QuestionExercisesResponse struct {
	TestTypeExerciseID string                   `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
	Content            string                   `json:"Content" form:"Content"`
	Weight             int                      `json:"Weight" form:"Weight"`
	Tags               []string                 `json:"Tags" form:"Tags"`
	ChoiceExercises    []ChoiceExercisesRequest `json:"ChoiceExercises"`
}

type ChoiceExercisesRequest struct {
	Type      string `json:"Type" form:"Type"`
	Content   string `json:"Content" form:"Content"`
	IsCorrect bool   `json:"IsCorrect" form:"IsCorrect"`
	Weight    int    `json:"Weight" form:"Weight"`
}
