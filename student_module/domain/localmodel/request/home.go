package request

type ParamMentorDetailHome struct {
	UUID string `uri:"uuid" validate:"required,uuid"`
}

type PlanHome struct {
	Module string `form:"module"`
	Name   string `form:"name"`
}
