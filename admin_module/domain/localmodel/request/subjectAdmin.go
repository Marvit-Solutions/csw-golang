package request

type ParamSubjectAdminAll struct {
	SearchKeywords string `form:"search_keywords"`
	SubModuleName  string `form:"sub_module_name"`
	Page           int    `form:"page"`
	Limit          int    `form:"limit"`
}

type ParamSubjectAdminDelete struct {
	UUID string `uri:"uuid"`
}

type ParamSubjectAdminUpdateDetail struct {
	UUID string `uri:"uuid"`
}

type SubjectAdminPayload struct {
	Name          string `json:"name"`
	SubModuleName string `json:"sub_module_name"`
}

type SubjectAdminPayloadUpdate struct {
	UUID          string `json:"uuid"`
	SubModuleName string `json:"sub_module_name"`
	Name          string `json:"name"`
}
