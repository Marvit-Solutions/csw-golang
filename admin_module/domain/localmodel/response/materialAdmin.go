package response

type MaterialAdminAllResponse struct {
	ID            int    `json:"id"`
	UUID          string `json:"uuid"`
	SubModuleName string `json:"sub_module_name"`
	ModuleName    string `json:"module_name"`
	SubjectName   string `json:"subject_name"`
	Name          string `json:"name"`
	Content       string `json:"content"`
}

type MaterialAdminUpdateDetailResponse struct {
	UUID        string `json:"uuid"`
	SubjectName string `json:"subject_name"`
	Name        string `json:"name"`
	Content     string `json:"content"`
}
