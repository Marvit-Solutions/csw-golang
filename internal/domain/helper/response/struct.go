package response

type ResponseError struct {
	Message string `json:"Message"`
	Code    int    `json:"Code"`
	Status  string `json:"Status"`
}

type ResponsePaged[Data interface{}] struct {
	Message string `json:"Message"`
	Code    int    `json:"Code"`
	Status  string `json:"Status"`
	Data    Data   `json:"Data"`
	Meta    *Meta  `json:"Meta"`
	// "Meta": {
	// 	"Page": "4"
	// 	"PerPage": "10",
	// 	"LastPage": "25",
	// 	"TotalPage": "50",
	//}
}

type ResponseNonPaged[Data interface{}] struct {
	Message string `json:"Message"`
	Code    int    `json:"Code"`
	Status  string `json:"Status"`
	Data    Data   `json:"Data"`
}

type Meta struct {
	Page      int64 `json:"Page"`
	PerPage   int64 `json:"PerPage"`
	LastPage  int64 `json:"LastPage"`
	TotalPage int64 `json:"TotalPage"`
}
