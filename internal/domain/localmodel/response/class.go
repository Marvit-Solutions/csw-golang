package response

import "github.com/Marvit-Solutions/csw-golang/library/struct/model"

type Class struct {
	UUID             string               `json:"uuid"`
	Module           string               `json:"module"`
	MeetingDate      string               `json:"meeting_date"`
	Media            *model.MultiResImage `json:"media"`
	Subject          SubjectClass         `json:"subject"`
	Mentor           MentorClass          `json:"mentor"`
	TotalParticipant int                  `json:"total_participant"`
	Duration         string               `json:"duration"`
}

type ClassScheduleList struct {
	UUID               string `json:"uuid"`
	ScheduleMediaID    *int   `json:"schedule_media_id"`
	ClassPlanTypeID    *int   `json:"class_plan_type_id"`
	ClassName          string `json:"class_name"`
	MentorUUID         string `json:"mentor_uuid"`
	ShortName          string `json:"short_name"`
	MentorMediaID      *int   `json:"mentor_media_id"`
	SubjectUUID        string `json:"subject_uuid"`
	SubjectName        string `json:"subject_name"`
	SubjectDescription string `json:"subject_description"`
	MeetingDate        string `json:"meeting_date"`
	MeetingLink        string `json:"meeting_link"`
	Duration           string `json:"duration"`
	ModuleName         string `json:"module_name"`
	ModuleSlug         string `json:"module_slug"`
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
