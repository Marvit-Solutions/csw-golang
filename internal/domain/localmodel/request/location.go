package request

type LocationRequest struct {
	Url        string
	ProvinceID string `uri:"province_id"`
	RegencyID  string `uri:"regency_id"`
	Province   string `form:"province"`
	Regency    string `form:"regency"`
	District   string `form:"district"`
}
