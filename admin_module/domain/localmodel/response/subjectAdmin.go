package response

type SubjectAdminAllResponse struct {
	UUID          string `json:"uuid"`
	SubModuleName string `json:"sub_module_name"`
	ModuleName    string `json:"module_name"`
	Name          string `json:"name"`
}

type SubjectAdminUpdateDetailResponse struct {
	UUID          string `json:"uuid"`
	SubModuleName string `json:"sub_module_name"`
	Name          string `json:"name"`
}
