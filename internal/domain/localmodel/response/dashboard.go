package response

import "time"

type MaterialDashboard struct {
	UUID      string    `json:"uuid"`
	Module    Module    `json:"module"`
	SubModule SubModule `json:"sub_module"`
	Subject   string    `json:"subject"`
}

type Module struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type SubModule struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type QuizAllDashboardPerTest struct {
	TestTypeName              string                       `json:"test_type"`
	QuizAllDashboardPerModule []*QuizAllDashboardPerModule `json:"quiz_all_dahsboard_per_module"`
}

type QuizAllDashboardPerModule struct {
	UUID       string                  `json:"uuid"` // uuid modul
	ModuleName string                  `json:"module_name"`
	Quizzes    []*QuizItemAllDashboard `json:"quizzes"`
}

type QuizItemAllDashboard struct {
	UUID               string    `json:"uuid"`
	SubModuleName      string    `json:"sub_module_name"`
	SubModuleUUID      string    `json:"sub_module_uuid"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	Time               int       `json:"time"`
	Open               time.Time `json:"open"`
	Close              time.Time `json:"close"`
	QuizSubmissionUUID string    `json:"quiz_submission_uuid"`
	MaxScore           int       `json:"max_score"`
	Attempt            int       `json:"attempt"`
	StatusPengerjaan   string    `json:"status_pengerjaan"`
}
