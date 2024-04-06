package response

type MentorHome struct {
	UUID           string `json:"UUID"`
	Name           string `json:"Name"`
	Type           string `json:"Type"`
	Description    string `json:"Description"`
	Motto          string `json:"Motto"`
	ProfilePicture string `json:"ProfilePicture"`
}

type MentorDetailHome struct {
	UUID           string `json:"UUID"`
	Name           string `json:"Name"`
	Type           string `json:"Type"`
	Description    string `json:"Description"`
	Motto          string `json:"Motto"`
	ProfilePicture string `json:"ProfilePicture"`
}

type PlanHome struct {
	UUID       string  `json:"UUID"`
	ModuleName string  `json:"ModuleName"`
	Name       string  `json:"Name"`
	Price      float64 `json:"Price"`
	Group      bool    `json:"Group"`
	Exercise   int     `json:"Exercise"`
	Access     int     `json:"Access"`
	Module     bool    `json:"Module"`
	TryOut     int     `json:"TryOut"`
	Zoom       bool    `json:"Zoom"`
}
