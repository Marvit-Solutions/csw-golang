package dto

type ModulResponse struct {
	ID          string `json:"ID" form:"ID"`
	Name        string `json:"Name" form:"Name"`
	Description string `json:"Description" form:"Description"`
	Subject     struct {
		ID   string `json:"ID" form:"ID"`
		Name string `json:"Name" form:"Name"`
	}
	Exercise struct {
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
	ID          string     `json:"ID" form:"ID"`
	Name        string     `json:"Name" form:"Name"`
	Description string     `json:"Description" form:"Description"`
	Explanation string     `json:"Explanation" form:"Explanation"`
	Status      string     `json:"Status" form:"Status"`
	Time        int        `json:"Time" form:"Time"`
	JumlahSoal  int        `json:"JumlahSoal" form:"JumlahSoal"`
	Question    []Question `json:"Soal" form:"Soal"`
}

type Question struct {
	ID       string   `json:"ID" form:"ID"`
	Number   int      `json:"Nomor" form:"Nomor"`
	Status   string   `json:"Status" form:"Status"`
	Mark     int      `json:"Mark" form:"Mark"`
	Flag     bool     `json:"Flag" form:"Flag"`
	Question string   `json:"Question" form:"Question"`
	Answers  []Answer `json:"Answers" form:"Answers"`
}

type Answer struct {
	ID      string  `json:"ID" form:"ID"`
	Jenis   string  `json:"Jenis" form:"Jenis"`
	Content string  `json:"Content" form:"Content"`
	Dipilih bool    `json:"Dipilih" form:"Dipilih"`
	Nilai   float64 `json:"Nilai" form:"Nilai"`
}

type ReviewResultResponse struct {
	ID            string     `json:"ID" form:"ID"`
	Name          string     `json:"Name" form:"Name"`
	Status        string     `json:"Status" form:"Status"`
	Start         string     `json:"Start" form:"Start"`
	Finish        string     `json:"Finish" form:"Finish"`
	Time          string     `json:"Time" form:"Time"`
	Mark          string     `json:"Mark" form:"Mark"`
	Grade         string     `json:"Grade" form:"Grade"`
	TimesFinished int        `json:"TimesFinished" form:"TimesFinished"`
	Question      []Question `json:"Soal" form:"Soal"`
}

type HistoryTop3NilaiReviewResponse struct {
	ModulID string `json:"ModulID" form:"ModulID"`
	Subject []struct {
		SubjectID string `json:"SubjectID" form:"SubjectID"`
		Name      string `json:"Name" form:"Name"`
		Nilai     []struct {
			ResultID string `json:"ResultID" form:"ResultID"`
			Mark     string `json:"Mark" form:"Mark"`
			Grade    string `json:"Grade" form:"Grade"`
		}
	}
}
