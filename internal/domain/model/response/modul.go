package response

type ModulResponse struct {
	UUID          string `json:"uuid"`
	SubModuleName string `json:"sub_module_name"`
	ModuleName    string `json:"module_name"`
	Description   string `json:"description"`
}
