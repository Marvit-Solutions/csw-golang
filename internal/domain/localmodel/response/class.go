package response

import "github.com/Marvit-Solutions/csw-golang/library/struct/model"

type Class struct {
	UUID         string               `json:"uuid"`
	Module       string               `json:"module"`
	MeetingDate  string               `json:"meeting_date"`
	Media        *model.MultiResImage `json:"media"`
	Subject      SubjectClass         `json:"subject"`
	Mentor       MentorClass          `json:"mentor"`
	TotalStudent int                  `json:"total_student"`
	Duration     string               `json:"duration"`
}

type MentorClass struct {
	UUID  string               `json:"uuid"`
	Name  string               `json:"name"`
	Media *model.MultiResImage `json:"media"`
}

type SubjectClass struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
