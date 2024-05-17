package request

type QuizContentType string

type ParamQuizContent struct {
	QuizUUID string `uri:"quiz_uuid"`
}

type ParamQuizReview struct {
	QuizSubmissionUUID string `uri:"quiz_submission_uuid"`
	QuizUUID           string `uri:"quiz_uuid"`
	TestTypeId         int    `uri:"test_type_id"`
}

type ParamQuizScoreAll struct {
	SubModulUUID string `uri:"sub_modul_uuid"`
}

type ParamQuizDetail struct {
	QuizUUID   string `uri:"quiz_uuid"`
	TestTypeId int    `uri:"test_type_id"`
}

type QuizSubmissionRequest struct {
	UserID         int                      `json:"user_id"`
	QuizID         int                      `json:"quiz_id"`
	TotalQuestions int                      `json:"total_questions"`
	TimeRequired   string                   `json:"time_required"`
	Questions      []QuestionSubmissionItem `json:"questions"`
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
