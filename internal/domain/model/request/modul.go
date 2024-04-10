package request

type ParamModulDetail struct {
	UUID string `uri:"uuid" validate:"required,uuid"`
}
