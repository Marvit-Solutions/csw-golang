package response

type LocationResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProvinceResponse struct {
	Code     string `json:"code"`
	Messages string `json:"messages"`
	Value    []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
}
