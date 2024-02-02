package datastruct

type Mentors struct {
	ID             string  `gorm:"type:text;primaryKey"`
	Name           string  `json:"Name" form:"Name"`
	Description    string  `json:"Description" form:"Description"`
	ProfilePicture string  `json:"ProfilePicture" form:"ProfilePicture"`
	Rating         float32 `json:"Rating" form:"Rating"`
	SubModuleID    string  `json:"SubModuleID" form:"SubModuleID"`
}
