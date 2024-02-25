package dto

import "time"

type QuizResponse struct {
	ID          string    `json:"ID"`
	TestType    string    `json:"TestType"`
	Title       string    `json:"Title"`
	MeetingDate time.Time `json:"MeetingDate"`
	Open        time.Time `json:"Open"`
	Description string    `json:"Description"`
	Time        uint      `json:"Time"`
	Point       uint      `json:"Point"`
	Attempt     uint      `json:"Attempt"`
}
