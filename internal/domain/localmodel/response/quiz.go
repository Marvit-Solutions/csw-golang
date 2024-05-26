package response

import (
	"time"
)

type QuestionItem struct {
	ID       int          `json:"id"`
	UUID     string       `json:"uuid"`
	Question string       `json:"question"`
	Options  []OptionItem `json:"options"`
	NoSoal   int          `json:"no_soal"`
	Status   string       `json:"status"`
	Mark     int          `json:"mark"`
}

type QuestionReviewItem struct {
	ID          int                `json:"id"`
	UUID        string             `json:"uuid"`
	Question    string             `json:"question"`
	Options     []OptionItemReview `json:"options"`
	NoSoal      int                `json:"no_soal"`
	Status      string             `json:"status"`
	Mark        int                `json:"mark"`
	Explanation string             `json:"explanation"`
}

type QuestionReviewItemStatus string

const (
	BelumDiJawab QuestionReviewItemStatus = "belum-dijawab"
	SudahDiJawab QuestionReviewItemStatus = "sudah-dijawab"
)

type QuestionsReview struct {
	QuestionReviewItem QuestionReviewItem `json:"question_review_item"`
	UserAnswer         int                `json:"user_answer"`  // nilainya merujuk ke id tabel quiz choice
	RightAnswer        int                `json:"right_answer"` // nilainya merujuk ke id tabel quiz choice
	RightAnswerText    string             `json:"right_answer_text"`
}

type OptionItem struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	Letter string `json:"letter"`
	Text   string `json:"text"`
}

type OptionItemSubmission struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type OptionItemReview struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	Text      string `json:"text"`
	Letter    string `json:"letter"`
	IsCorrect bool   `json:"is_correct"`
}

type QuizContentResponse struct {
	ID             int            `json:"id"`
	UUID           string         `json:"uuid"`
	Topic          string         `json:"topic"`
	Modul          string         `json:"modul"`
	Description    string         `json:"description"`
	TotalQuestions int            `json:"total_questions"`
	TotalTime      int            `json:"total_time"`
	Questions      []QuestionItem `json:"questions"`
}

type QuizReviewResponse struct {
	ID                int               `json:"id"`
	UUID              string            `json:"uuid"`
	Topic             string            `json:"topic"`
	Modul             string            `json:"modul"`
	TotalQuestions    int               `json:"total_questions"`
	StartedAt         string            `json:"started_at"`
	TotalTime         string            `json:"total_time"`
	TotalRightAnswers int               `json:"total_right_answers"`
	Score             int               `json:"score"`
	MaxScore          int               `json:"max_score"`
	Attempt           int               `json:"attempt"`
	Questions         []QuestionsReview `json:"questions"`
}

type QuizStatus string

const (
	BelumDikerjakan QuizStatus = "belum-dikerjakan"
	SudahDikerjakan QuizStatus = "sudah-dikerjakan"
)

type QuizDetailResponse struct {
	ID                 int        `json:"id"`
	UUID               string     `json:"uuid"`
	Title              string     `json:"title"`
	Subject            string     `json:"subject"`
	Modul              string     `json:"modul"`
	Description        string     `json:"description"`
	TotalQuestions     int        `json:"total_questions"`
	TotalTime          int        `json:"total_time"`
	Status             QuizStatus `json:"status,omitempty"`
	AttemptAllowed     int        `json:"attempt_allowed"`
	QuizSubmissionUUID string     `json:"quiz_submission_uuid,omitempty"`
	Score              int        `json:"score"`
	MaxScore           int        `json:"max_score"`
	Attempt            int        `json:"attempt"`
}

type UserAnswer struct {
	QuestionId      int    `json:"question_id"`
	QuestionContent string `json:"question_content"`
	ChoiceId        int    `json:"choice_id"`
	ChoiceContent   string `json:"choice_content"`
}

type QuizScoreAllResponse struct {
	ID               int                              `json:"id"` // id submodul
	SubModul         string                           `json:"sub_modul"`
	QuizScoreAllItem []QuizScoreAllItemGroupBySubject `json:"quiz_score_all_item"`
}

type QuizScoreAllItemTmp struct {
	SubjectID          int    `json:"subject_id"`
	Subject            string `json:"subject"`
	QuizUUID           string `json:"quiz_uuid"`
	Quiz               string `json:"quiz"`
	QuizSubmissionUUID string `json:"quiz_submission_uuid,omitempty"`
	Score              int    `json:"score,omitempty"`
	MaxScore           int    `json:"max_score,omitempty"`
	TotalRightAnswers  int    `json:"total_right_answers"`
	TotalQuestions     int    `json:"total_questions"`
}

type QuizScoreAllItemGroupBySubject struct {
	SubjectID        int                       `json:"subject_id"`
	Subject          string                    `json:"subject"`
	QuizScoresDetail []QuizScoreAllSubjectItem `json:"quiz_scores_detail"`
}

type QuizScoreAllSubjectItem struct {
	QuizSubmissionUUID string `json:"quiz_submission_uuid"`
	QuizUUID           string `json:"quiz_uuid"`
	Quiz               string `json:"quiz"`
	Score              int    `json:"score"`
	MaxScore           int    `json:"max_score"`
	TotalRightAnswers  int    `json:"total_right_answers"`
	TotalQuestions     int    `json:"total_questions"`
}

type QuizSubModuleAllResponse struct {
	UUID                      string                       `json:"uuid"`
	ModuleName                string                       `json:"module_name"`
	QuizzesGroupedBySubModule []*QuizzesGroupedBySubModule `json:"quizzes_grouped_by_sub_module"`
}

type QuizzesGroupedBySubModule struct {
	SubModuleUUID   string `json:"sub_module_uuid"`
	SubModuleID     int    `json:"sub_module_id"`
	SubModuleName   string `json:"sub_module_name"`
	QuizCount       int    `json:"quiz_count"`
	SubmissionCount int    `json:"submission_count"`
}

type QuizAllResponse struct {
	UUID          string         `json:"uuid"`
	SubModuleName string         `json:"sub_module_name"`
	ModuleName    string         `json:"module_name"`
	Quizzes       []*QuizItemAll `json:"quizzes"`
}

type QuizItemAll struct {
	ID                 int       `json:"id"`
	UUID               string    `json:"uuid"`
	Subject            string    `json:"sub_subject"`
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
