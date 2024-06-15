package response

import (
	"time"
)

type QuizAdminAllResponse struct {
	ID            int       `json:"id"`
	UUID          string    `json:"uuid"`
	SubModuleName string    `json:"sub_module_name"`
	ModuleName    string    `json:"module_name"`
	SubjectName   string    `json:"subject_name"`
	TestTypeName  string    `json:"test_type_name"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Time          int       `json:"time"`
	Open          time.Time `json:"open"`
	Close         time.Time `json:"close"`
	MaxScore      int       `json:"max_score"`
	Attempt       int       `json:"attempt"`
}

type ChoiceQuizAdmin struct {
	UUID      string `json:"uuid"`
	Content   string `json:"content"`
	Score     int    `json:"score"`
	IsCorrect bool   `json:"is_correct"`
}

type QuestionQuizAdmin struct {
	ID          int               `json:"id"`
	UUID        string            `json:"uuid"`
	Content     string            `json:"content"`
	Score       int               `json:"score"`
	Explanation string            `json:"explanation"`
	Choices     []ChoiceQuizAdmin `json:"choices"`
}

type QuizAdminUpdateDetailResponse struct {
	UUID        string              `json:"uuid"`
	Title       string              `json:"title"`
	Subject     string              `json:"subject"`
	TestType    string              `json:"test_type"`
	Time        int                 `json:"time"`
	Attempt     int                 `json:"attempt"`
	Open        time.Time           `json:"open"`
	Close       time.Time           `json:"close"`
	Description string              `json:"description"`
	Questions   []QuestionQuizAdmin `json:"questions"`
}
