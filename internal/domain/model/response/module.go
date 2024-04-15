package response

type ModulResponse struct {
	UUID          string `json:"uuid"`
	SubModuleName string `json:"sub_module_name"`
	ModuleName    string `json:"module_name"`
	Description   string `json:"description"`
}

type ModuleDetailResponse struct {
	UUID          string             `json:"uuid"`
	SubModuleName string             `json:"sub_module_name"`
	ModuleName    string             `json:"module_name"`
	Description   string             `json:"description"`
	Subjects      []*SubjectResponse `json:"subjects"`
	Quizzes       []*QuizResponse    `json:"quizzes"`
}

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

type MaterialResponse struct {
	UUID       string                `json:"uuid"`
	Module     string                `json:"module"`
	SubModule  string                `json:"sub_module"`
	Subject    string                `json:"subject"`
	SubSubject []*SubSubjectResponse `json:"sub_subject"`
}

type SubSubjectResponse struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type MaterialFindResponse struct {
	Subjects    []*SubjectFindResponse    `json:"subjects"`
	SubSubjects []*SubSubjectFindResponse `json:"sub_subjects"`
}

type SubjectFindResponse struct {
	ID            int     `json:"-"`
	UUID          string  `json:"uuid"`
	Name          string  `json:"name"`
	Module        string  `json:"module"`
	SubModule     string  `json:"sub_module"`
	CountMaterial int     `json:"count_material"`
	LastUpdated   *string `json:"last_updated"`
}

type SubSubjectFindResponse struct {
	ID          int     `json:"-"`
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Module      string  `json:"module"`
	SubModule   string  `json:"sub_module"`
	LastUpdated *string `json:"last_updated"`
}
