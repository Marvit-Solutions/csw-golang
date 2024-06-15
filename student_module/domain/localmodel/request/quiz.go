package request

type QuizContentType string

type ParamQuizContent struct {
	QuizUUID string `uri:"quiz_uuid"`
}

type ParamQuizReview struct {
	AuthenticatedUser  int    `json:"authenticated_user"`
	QuizSubmissionUUID string `uri:"quiz_submission_uuid"`
	TestTypeId         int    `uri:"test_type_id"`
}

type ParamQuizScoreAll struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	SubModulUUID      string `uri:"sub_modul_uuid"`
}

type ParamQuizDetail struct {
	AuthenticatedUser int    `json:"authenticated_user"`
	QuizUUID          string `uri:"quiz_uuid"`
	TestTypeId        int    `uri:"test_type_id"`
}

type ParamQuizSubModuleAll struct {
	ModuleID   int `uri:"module_id"`
	TestTypeId int `uri:"test_type_id"`
}

type TestStatus string

const (
	BelumDiJawab TestStatus = "belum-dijawab"
	SudahDiJawab TestStatus = "sudah-dijawab"
	All          TestStatus = "all"
)

type ParamQuizAll struct {
	AuthenticatedUser int        `json:"authenticated_user"`
	SubModuleUUID     string     `uri:"sub_module_uuid"`
	TestTypeId        int        `uri:"test_type_id"`
	TestStatus        TestStatus `uri:"test_status"`
	Page              int        `form:"page"`
	Limit             int        `form:"limit"`
}

type QuizSubmissionRequest struct {
	AuthenticatedUser int                      `json:"authenticated_user"`
	QuizID            int                      `json:"quiz_id"`
	TotalQuestions    int                      `json:"total_questions"`
	TimeRequired      string                   `json:"time_required"`
	Questions         []QuestionSubmissionItem `json:"questions"`
}

type QuestionSubmissionItem struct {
	ID         int          `json:"id"`
	UUID       string       `json:"uuid"`
	Question   string       `json:"question"`
	Options    []OptionItem `json:"options"`
	Status     string       `json:"status"`
	UserAnswer int          `json:"user_answer"`
	Mark       int          `json:"mark"`
}

type OptionItem struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	Letter string `json:"letter"`
	Text   string `json:"text"`
}
