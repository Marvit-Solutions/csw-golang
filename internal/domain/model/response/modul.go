package response

type ModulResponse struct {
	UUID          string `json:"uuid"`
	SubModuleName string `json:"sub_module_name"`
	ModuleName    string `json:"module_name"`
	Description   string `json:"description"`
}

type ModulDetailResponse struct {
	UUID          string             `json:"uuid"`
	SubModuleName string             `json:"sub_module_name"`
	ModuleName    string             `json:"module_name"`
	Description   string             `json:"description"`
	Subjects      []*SubjectResponse `json:"subjects"`
	Quizzes       []*QuizResponse    `json:"quizzes"`
}

// type MaterialResponse struct {
// 	Subjects []*SubjectResponse `json:"subjects"`
// 	Quizzes  []*QuizResponse    `json:"quizzes"`
// }

type SubjectResponse struct {
	ID   int    `json:"-"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type QuizResponse struct {
	ID    int    `json:"-"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
}
