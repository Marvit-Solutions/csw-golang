package request

type LocationRequest struct {
	Url        string
	ProvinceID string `uri:"province"`
	RegencyID  string `uri:"regency"`
	Province   string `form:"province"`
	Regency    string `form:"regency"`
	District   string `form:"district"`
}
