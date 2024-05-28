package response

type MaterialDashboard struct {
	UUID      string    `json:"uuid"`
	Module    Module    `json:"module"`
	SubModule SubModule `json:"sub_module"`
	Subject   string    `json:"subject"`
}

type Module struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type SubModule struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
