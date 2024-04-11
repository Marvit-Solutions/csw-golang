package request

type ParamModuleDetail struct {
	UUID string `uri:"uuid" validate:"required,uuid"`
}

type ParamMaterial struct {
	SubjectUUID string `uri:"subject_uuid" validate:"required,uuid"`
}
