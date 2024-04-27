package request

type ParamModule struct {
	SubModuleUUID string `uri:"sub_module_uuid"`
	SubjectUUID   string `uri:"subject_uuid"`
}

type Material struct {
	Subject    bool `form:"subject" validate:"required"`
	SubSubject bool `form:"sub_subject" validate:"required"`
}
