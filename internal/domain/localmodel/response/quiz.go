package response

type QuizContentResponse struct {
	ID             int            `json:"id"`
	UUID           string         `json:"uuid"`
	Topic          string         `json:"topic"`
	Modul          string         `json:"modul"`
	TotalQuestions int            `json:"total_questions"`
	TotalTime      int            `json:"total_time"`
	Questions      []QuestionItem `json:"questions"`
}

type QuestionItem struct {
	ID       int          `json:"id"`
	UUID     string       `json:"uuid"`
	Question string       `json:"question"`
	Options  []OptionItem `json:"options"`
	NoSoal   int          `json:"no_soal"`
	Status   string       `json:"status"`
	Mark     int          `json:"mark"`
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

// Get Quiz Detail api
type Status string

const (
	BelumDikerjakan Status = "belum-dikerjakan"
	SudahDikerjakan Status = "sudah-dikerjakan"
)

type AttemptAllowed string

const (
	BerkaliKali AttemptAllowed = "berkali-kali"
	Sekali      AttemptAllowed = "sekali"
)

type DetailQuizResponse struct {
	ID                 int            `json:"id"`
	UUID               string         `json:"uuid"`
	Subject            string         `json:"sub_subject"`
	Modul              string         `json:"modul"`
	TotalQuestions     int            `json:"total_questions"`
	TotalTime          int            `json:"total_time,omitempty"`
	Status             Status         `json:"status,omitempty"`
	AttemptAllowed     AttemptAllowed `json:"attempt_allowed,omitempty"`
	QuizSubmissionUUID string         `json:"quiz_submission_uuid,omitempty"`
	Score              string         `json:"score,omitempty"`
	Attempt            int            `json:"attempt,omitempty"`
}
