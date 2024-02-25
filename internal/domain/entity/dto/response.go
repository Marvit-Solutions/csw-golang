package dto

type Success[Data interface{}] struct {
	Message string `json:"Message"`
	Code    uint   `json:"Code"`
	Status  string `json:"Status"`
	Data    Data   `json:"Data,omitempty"`
	Meta    *Meta  `json:"Meta,omitempty"`
	// "Meta": {
	// 	"Page": "4"
	// 	"PerPage": "10",
	// 	"LastPage": "25",
	// 	"TotalPage": "50",
	//}
}

type Fail struct {
	Message string `json:"Message"`
	Code    uint   `json:"Code"`
	Status  string `json:"Status"`
}

type Meta struct {
	Page      int64 `json:"Page"`
	PerPage   int64 `json:"PerPage"`
	LastPage  int64 `json:"LastPage"`
	TotalPage int64 `json:"TotalPage"`
}
