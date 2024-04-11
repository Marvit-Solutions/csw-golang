package request

type ParamModuleDetail struct {
	UUID string `uri:"uuid" validate:"required,uuid"`
}
