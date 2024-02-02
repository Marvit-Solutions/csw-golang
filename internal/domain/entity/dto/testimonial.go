package dto

type Testimonials []struct {
	ID      string  `json:"ID"`
	Comment string  `json:"Comment"`
	Rating  float32 `json:"Rating"`
	User    struct {
		Name           string `json:"Name" form:"Name"`
		ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
		Status         string `json:"Status" form:"Status"`
	}
}
