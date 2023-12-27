package dto

type Success[Data interface{}] struct {
	Status  string `json:"Status"`
	Code    uint   `json:"Code"`
	Message string `json:"Message"`
	Data    Data   `json:"Data"`
}

type Fail struct {
	Code    uint   `json:"Code"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
}
