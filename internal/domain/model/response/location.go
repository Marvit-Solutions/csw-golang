package response

type LocationResponse struct {
	ID   string `json:"id"`
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

type DistrictResponse struct {
	Code     string `json:"code"`
	Messages string `json:"messages"`
	Value    []struct {
		ID        string `json:"id"`
		RegencyID string `json:"id_kabupaten"`
		Name      string `json:"name"`
	}
}
