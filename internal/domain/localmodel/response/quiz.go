package response

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

// Get Quiz Detail api
type QuizStatus string

const (
	BelumDikerjakan QuizStatus = "belum-dikerjakan"
	SudahDikerjakan QuizStatus = "sudah-dikerjakan"
)

type QuizDetailResponse struct {
	ID                 int        `json:"id"`
	UUID               string     `json:"uuid"`
	Subject            string     `json:"sub_subject"`
	Modul              string     `json:"modul"`
	Description        string     `json:"description"`
	TotalQuestions     int        `json:"total_questions"`
	TotalTime          int        `json:"total_time,omitempty"`
	Status             QuizStatus `json:"status,omitempty"`
	AttemptAllowed     int        `json:"attempt_allowed,omitempty"`
	QuizSubmissionUUID string     `json:"quiz_submission_uuid,omitempty"`
	Score              int        `json:"score,omitempty"`
	ScoreMax           int        `json:"score_max,omitempty"`
	Attempt            int        `json:"attempt"`
}

type UserAnswer struct {
	QuestionId      int    `json:"question_id"`
	QuestionContent string `json:"question_content"`
	ChoiceId        int    `json:"choice_id"`
	ChoiceContent   string `json:"choice_content"`
}
