package dto

type ModuleResponse struct {
	ID          string `json:"ID" form:"ID"`
	Name        string `json:"Name" form:"Name"`
	Description string `json:"Description" form:"Description"`
	Subject     []struct {
		ID   string `json:"ID" form:"ID"`
		Name string `json:"Name" form:"Name"`
	}
	Exercise []struct {
		ID   string `json:"ID" form:"ID"`
		Name string `json:"Name" form:"Name"`
	}
}

type SubjectResponse struct {
	ID         string       `json:"ID" form:"ID"`
	Name       string       `json:"Name" form:"Name"`
	SubSubject []SubSubject `json:"SubSubject" form:"SubSubject"`
}

type SubSubject struct {
	ID      string `json:"ID" form:"ID"`
	Name    string `json:"Name" form:"Name"`
	Content string `json:"Content" form:"Content"`
}

type ExerciseResponse struct {
	ID            string     `json:"ID" form:"ID"`
	Name          string     `json:"Name" form:"Name"`
	Description   string     `json:"Description" form:"Description"`
	Explanation   string     `json:"Explanation" form:"Explanation"`
	Status        string     `json:"Status" form:"Status"`
	Time          uint       `json:"Time" form:"Time"`
	QuestionTotal int        `json:"QuestionTotal" form:"QuestionTotal"`
	Questions     []Question `json:"Questions" form:"Questions"`
}

type Question struct {
	ID       string   `json:"ID" form:"ID"`
	Number   int      `json:"Number" form:"Number"`
	Status   string   `json:"Status" form:"Status"`
	Mark     int      `json:"Mark" form:"Mark"`
	Flag     bool     `json:"Flag" form:"Flag"`
	Question string   `json:"Question" form:"Question"`
	Image    string   `json:"Image" form:"Image"`
	Answers  []Answer `json:"Answers" form:"Answers"`
}

type Answer struct {
	ID        string `json:"ID" form:"ID"`
	Content   string `json:"Content" form:"Content"`
	IsCorrect bool   `json:"IsCorrect" form:"IsCorrect"`
	Weight    int    `json:"Weight" form:"Weight"`
	IsChosen  bool   `json:"IsChosen" form:"IsChosen"`
}

type ReviewResultResponse struct {
	ID            string     `json:"ID" form:"ID"`
	Name          string     `json:"Name" form:"Name"`
	Description   string     `json:"Description" form:"Description"`
	Explanation   string     `json:"Explanation" form:"Explanation"`
	Status        string     `json:"Status" form:"Status"`
	Start         string     `json:"Start" form:"Start"`
	Finish        string     `json:"Finish" form:"Finish"`
	Time          string     `json:"Time" form:"Time"`
	QuestionTotal int        `json:"QuestionTotal" form:"QuestionTotal"`
	Grade         int        `json:"Grade" form:"Grade"`
	TimesFinished int        `json:"TimesFinished" form:"TimesFinished"`
	Questions     []Question `json:"Soal" form:"Soal"`
}



type HistoryTop3ScoreResponse struct {
	SubModuleID   string `json:"SubModuleID" form:"SubModuleID"`
	SubModuleName string `json:"SubModuleName" form:"SubModuleName"`
	Subject       []HistoryTop3ScoreSubjectResponse
}

type HistoryTop3ScoreSubjectResponse struct {
	SubjectID string                          `json:"SubjectID" form:"SubjectID"`
	Name      string                          `json:"Name" form:"Name"`
	Grade     []HistoryTop3ScoreGradeResponse `json:"Grade" form:"Grade"`
}

type HistoryTop3ScoreGradeResponse struct {
	ResultID    string `json:"ResultID" form:"ResultID"`
	GradingTime string `json:"GradingTime" form:"GradingTime"`
	Mark        int    `json:"Mark" form:"Mark"`
	Score       int    `json:"Score" form:"Score"`
}
