package response

import "github.com/Marvit-Solutions/csw-golang/library/struct/model"

type ExerciseResponse struct {
	UUID          string `json:"uuid"`
	TestType      string `json:"test_type"`
	ModuleName    string `json:"module_name"`
	Title         string `json:"title"`
	TotalQuestion int    `json:"total_question"`
	Attempt       int    `json:"attempt"`
	Time          int    `json:"time"`
	Description   string `json:"description"`
}

type ExerciseDetailResponse struct {
	UUID        string      `json:"uuid"`
	TestType    string      `json:"test_type"`
	ModuleName  string      `json:"module_name"`
	Title       string      `json:"title"`
	Attempt     int         `json:"attempt"`
	UserAttempt int         `json:"user_attempt"`
	Time        int         `json:"time"`
	Description string      `json:"description"`
	Questions   []*Question `json:"questions"`
}

type Question struct {
	UUID          string           `json:"uuid"`
	Content       string           `json:"content"`
	QuestionMedia []*QuestionMedia `json:"question_medias"`
	Choices       []*Choice        `json:"choices"`
	Score         int              `json:"score"`
}

type QuestionMedia struct {
	Index int                  `json:"index"`
	Media *model.MultiResImage `json:"media"`
}

type Choice struct {
	UUID       string `json:"uuid"`
	QuestionID int    `json:"-"`
	Content    string `json:"content"`
}
