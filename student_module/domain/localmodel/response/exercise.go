package response

import "github.com/Marvit-Solutions/csw-golang/library/struct/model"

type Exercise struct {
	UUID          string `json:"uuid"`
	TestType      string `json:"test_type"`
	ModuleName    string `json:"module_name"`
	Title         string `json:"title"`
	TotalQuestion int    `json:"total_question"`
	Attempt       int    `json:"attempt"`
	Time          int    `json:"time"`
	Description   string `json:"description"`
}

type ExerciseDetail struct {
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

type ExerciseHistoryDetailSubModule struct {
	SubmissionUUID               string `json:"submission_uuid"`
	SubModule                    string `json:"submodule"`
	ScorePerSubModule            int    `json:"score_per_sub_module"`
	MaxScorePerSubModule         int    `json:"max_score_per_sub_module"`
	TotalRightAnswerPerSubModule int    `json:"total_right_answer_per_sub_module"`
	TotalQuestionPerSubModule    int    `json:"total_question_per_sub_module"`
}

type ExerciseHistoryDetailModule struct {
	UUID                   string                            `json:"uuid"`
	ScorePerModule         int                               `json:"score_per_module"`
	MaxScorePerModule      int                               `json:"max_score_per_module"`
	RightAnswerPerModule   int                               `json:"right_answer_per_module"`
	TotalQuestionPerModule int                               `json:"total_question_per_module"`
	ExerciseHistoryDetail  []*ExerciseHistoryDetailSubModule `json:"exercise_history_detail"`
}

type ExerciseHistory struct {
	MaxScore *ExerciseHistoryDetailModule `json:"max_score"`
	MinScore *ExerciseHistoryDetailModule `json:"min_score"`
}

type Question struct {
	ID            int              `json:"id"`
	UUID          string           `json:"uuid"`
	SubModuleID   int              `json:"sub_module_id"`
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
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	QuestionID int    `json:"-"`
	Content    string `json:"content"`
}

type ExerciseReview struct {
	ID            int               `json:"id"`
	UUID          string            `json:"uuid"`
	Topic         string            `json:"topic"`
	Modul         string            `json:"modul"`
	StartedAt     string            `json:"started_at"`
	FinishedAt    string            `json:"finished_at"`
	TimeRequired  string            `json:"total_time"`
	Attempt       int               `json:"attempt"`
	RightAnswer   int               `json:"total_right_answers"`
	TotalQuestion int               `json:"total_questions"`
	Score         int               `json:"score"`
	PerfectScore  int               `json:"max_score"`
	Questions     []*QuestionReview `json:"questions"`
}

type QuestionReview struct {
	ID              int              `json:"id"`
	UUID            string           `json:"uuid"`
	Content         string           `json:"question"`
	Explanation     string           `json:"explanation"`
	Score           int              `json:"mark"`
	Status          string           `json:"status"`
	UserAnswer      int              `json:"user_answer"`
	RightAnswer     int              `json:"right_answer"`
	RightAnswerText string           `json:"right_answer_text"`
	QuestionMedia   []*QuestionMedia `json:"question_medias"`
	Choices         []*ChoiceReview  `json:"options"`
}

type ChoiceReview struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	QuestionID int    `json:"-"`
	Content    string `json:"text"`
	IsChoose   bool   `json:"is_choose"`
	IsCorrect  bool   `json:"is_correct"`
}
