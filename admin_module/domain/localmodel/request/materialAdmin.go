package request

type ParamMaterialAdminAll struct {
	SearchKeywords string `form:"search_keywords"`
	SubModuleName  string `form:"sub_module_name"`
	SubjectName    string `form:"subject_name"`
	Page           int    `form:"page"`
	Limit          int    `form:"limit"`
}

type ParamMaterialAdminDelete struct {
	UUID string `uri:"uuid"`
}

type ParamMaterialAdminUpdateDetail struct {
	UUID string `uri:"uuid"`
}

type MaterialAdminPayload struct {
	SubjectName string `json:"subject_name"`
	Name        string `json:"name"`
	Content     string `json:"content"`
}

type MaterialAdminPayloadUpdate struct {
	UUID        string `json:"uuid"`
	SubjectName string `json:"subject_name"`
	Name        string `json:"name"`
	Content     string `json:"content"`
}
