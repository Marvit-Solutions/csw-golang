package request

type ParamExercise struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	Module            string `form:"module"`
	TestType          string `form:"test_type"`
}

type ExerciseDetailRequest struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	ExerciseUUID      string `uri:"exercise_uuid"`
}

type ExerciseCreateRequest struct {
	AuthenticatedUser int      `json:"authenticated_user"`
	ExerciseUUID      string   `json:"exercise_uuid"`
	TimeRequired      string   `json:"time_required"`
	Answers           []Answer `json:"answers"`
}

type ExerciseHistory struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	ExerciseUUID      string `uri:"exercise_uuid"`
}

type Answer struct {
	ChoiceUUID string `json:"choice_uuid"`
}

type ExerciseReview struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	SubmissionUUID    string `uri:"submission_uuid"`
}
