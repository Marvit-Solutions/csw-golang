package dto

type Success[Data interface{}] struct {
	Status  string `json:"Status"`
	Code    uint   `json:"Code"`
	Message string `json:"Message"`
	Data    Data   `json:"Data"`
}

type Fail struct {
	Status  string `json:"Status"`
	Code    uint   `json:"Code"`
	Message string `json:"Message"`
}
