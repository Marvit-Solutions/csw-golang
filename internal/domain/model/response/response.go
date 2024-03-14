package response

type Success[Data interface{}] struct {
	Message string      `json:"Message"`
	Code    uint        `json:"Code"`
	Status  string      `json:"Status"`
	Data    Data        `json:"Data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
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
