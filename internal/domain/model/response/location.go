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

type RegencyResponse struct {
	Code     string `json:"code"`
	Messages string `json:"messages"`
	Value    []struct {
		ID         string `json:"id"`
		ProvinceID string `json:"id_provinsi"`
		Name       string `json:"name"`
	}
}
