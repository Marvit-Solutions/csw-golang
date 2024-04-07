package request

type ParamMentorDetailHome struct {
	UUID string `uri:"uuid" validate:"required,uuid"`
}
