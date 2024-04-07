package response

type MentorHome struct {
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	TeachingField  string `json:"teaching_field"`
	Description    string `json:"description"`
	Motto          string `json:"motto"`
	ProfilePicture string `json:"profile_picture"`
}

type MentorDetailHome struct {
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	TeachingField  string `json:"teaching_field"`
	Description    string `json:"description"`
	Motto          string `json:"motto"`
	ProfilePicture string `json:"profile_picture"`
}

type PlanHome struct {
	UUID       string  `json:"uuid"`
	ModuleName string  `json:"module_name"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Group      bool    `json:"group"`
	Exercise   int     `json:"exercise"`
	Access     int     `json:"access"`
	Module     bool    `json:"module"`
	TryOut     int     `json:"try_out"`
	Zoom       bool    `json:"zoom"`
}
