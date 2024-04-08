package request

type LocationRequest struct {
	Url      string
	Province string `form:"province"`
	Regency  string `form:"regency"`
	District string `form:"district"`
}
