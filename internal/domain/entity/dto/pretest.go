package dto

type GetAllPretestResponse struct {
	ID        string              `json:"ID" form:"ID"`
	Name      string              `json:"Name" form:"Name"`
	SubModule []*SubModulePretest `json:"SubModule" form:"SubModule"`
}

type SubModulePretest struct {
	IDSubModule          string            `json:"IDSubModule" form:"IDSubModule"`
	NameSubModule        string            `json:"NameSubModule" form:"NameSubModule"`
	DescriptionSubModule string            `json:"DescriptionSubModule" form:"DescriptionSubModule"`
	Subject              []*SubjectPretest `json:"Subject,omitempty" form:"Subject,omitempty"`
}

type SubjectPretest struct {
	IDSubject   string     `json:"IDSubject" form:"IDSubject"`
	NameSubject string     `json:"NameSubject" form:"NameSubject"`
	Pretest     []*Pretest `json:"Pretest,omitempty" form:"Pretest,omitempty"`
}

type Pretest struct {
	IDPretest   string `json:"IDPretest" form:"IDPretest"`
	TestType    string `json:"TestType" form:"TestType"`
	Title       string `json:"Title" form:"Title"`
	MeetingDate string `json:"MeetingDate" form:"MeetingDate"`
	Open        string `json:"Open" form:"Open"`
	Description string `json:"Description" form:"Description"`
	Time        uint   `json:"Time" form:"Time"`
	Point       int    `json:"Point" form:"Point"`
	Status      string `json:"Status" form:"Status"`
	Attempt     int    `json:"Attempt" form:"Attempt"`
}
