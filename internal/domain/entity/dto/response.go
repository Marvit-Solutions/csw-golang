package dto

type Success[Data interface{}] struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Fail struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
