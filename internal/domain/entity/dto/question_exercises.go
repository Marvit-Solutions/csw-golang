package dto

type QuestionExercisesRequest struct {
	Content         string                   `json:"Content" form:"Content"`
	Weight          int                      `json:"Weight" form:"Weight"`
	ChoiceExercises []ChoiceExercisesRequest `json:"ChoiceExercises"`
}

type QuestionExercisesResponse struct {
	Content         string                   `json:"Content" form:"Content"`
	Weight          int                      `json:"Weight" form:"Weight"`
	ChoiceExercises []ChoiceExercisesRequest `json:"ChoiceExercises"`
}

type ChoiceExercisesRequest struct {
	Type      string `json:"Type" form:"Type"`
	Content   string `json:"Content" form:"Content"`
	IsCorrect bool   `json:"IsCorrect" form:"IsCorrect"`
	Weight    int    `json:"Weight" form:"Weight"`
}
