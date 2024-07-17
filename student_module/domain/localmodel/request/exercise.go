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

type ExerciseOptionItem struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
}

type QuestionCreateItem struct {
	ID          int                  `json:"id"`
	UUID        string               `json:"uuid"`
	SubModuleID int                  `json:"sub_module_id"`
	Options     []ExerciseOptionItem `json:"options"`
	UserAnswer  string               `json:"user_answer"`
	Score       int                  `json:"score"`
}
type ExerciseCreateRequest struct {
	AuthenticatedUser int                  `json:"authenticated_user"`
	ExerciseUUID      string               `json:"exercise_uuid"`
	TimeRequired      string               `json:"time_required"`
	Questions         []QuestionCreateItem `json:"questions"`
}

type ExerciseHistory struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	ExerciseUUID      string `uri:"exercise_uuid"`
}

type Answer struct {
	SubModuleID int    `json:"sub_module_id"`
	ChoiceUUID  string `json:"choice_uuid"`
}

type ExerciseReview struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	SubmissionUUID    string `uri:"submission_uuid"`
}
