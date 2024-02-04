package dto

type GetAllPretestResponse struct {
	ID      string `json:"ID" form:"ID"`
	Name    string `json:"Name" form:"Name"`
	SubPlan struct {
		ID          string `json:"ID" form:"ID"`
		Name        string `json:"Name" form:"Name"`
		Description string `json:"Description" form:"Description"`
		Pretest     []struct {
			ID          string `json:"ID" form:"ID"`
			Title       string `json:"Title" form:"Title"`
			MeetingDate string `json:"MeetingDate" form:"MeetingDate"`
			Open        string `json:"Open" form:"Open"`
			Description string `json:"Description" form:"Description"`
			Time        string `json:"Time" form:"Time"`
			Point       int    `json:"Point" form:"Point"`
			Status      string `json:"Status" form:"Status"`
			Attempt     int    `json:"Attempt" form:"Attempt"`
		}
	}
}
