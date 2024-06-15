package request

type ParamQuizAdminAll struct {
	SearchKeywords string `form:"search_keywords"`
	SubModuleName  string `form:"sub_module_name"`
	TestTypeName   string `form:"test_type_name"`
	Page           int    `form:"page"`
	Limit          int    `form:"limit"`
}

type ParamQuizAdminDelete struct {
	UUID string `uri:"uuid"`
}

type ParamQuizAdminUpdateDetail struct {
	UUID string `uri:"uuid"` // quiz uuid
}
type Choice struct {
	Content   string `json:"content"`
	Score     int    `json:"score"`
	IsCorrect bool   `json:"is_correct"`
}

type Question struct {
	Content     string   `json:"content"`
	Score       int      `json:"score"`
	Explanation string   `json:"explanation"`
	Choices     []Choice `json:"choices"`
}

type QuizAdminPayload struct {
	Title       string     `json:"title"`
	Subject     string     `json:"subject"`
	TestType    string     `json:"test_type"`
	Time        int        `json:"time"`
	Attempt     int        `json:"attempt"`
	Open        string     `json:"open"`
	Close       string     `json:"closed"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type ChoiceUpdate struct {
	UUID      string `json:"uuid"`
	Content   string `json:"content"`
	Score     int    `json:"score"`
	IsCorrect bool   `json:"is_correct"`
}

type QuestionUpdate struct {
	UUID        string         `json:"uuid"`
	Content     string         `json:"content"`
	Score       int            `json:"score"`
	Explanation string         `json:"explanation"`
	Choices     []ChoiceUpdate `json:"choices"`
}

type QuizAdminPayloadUpdate struct {
	UUID        string           `json:"uuid"`
	Title       string           `json:"title"`
	Subject     string           `json:"subject"`
	TestType    string           `json:"test_type"`
	Time        int              `json:"time"`
	Attempt     int              `json:"attempt"`
	Open        string           `json:"open"`
	Close       string           `json:"closed"`
	Description string           `json:"description"`
	Questions   []QuestionUpdate `json:"questions"`
}
