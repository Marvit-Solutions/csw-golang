package request

type QuizParamRequest struct {
	UserID    string `json:"UserID" form:"UserID"`
	Module    string `json:"Module" form:"Module"`
	SubModule string `json:"SubModule" form:"SubModule"`
	TestType  string `json:"TestType" form:"TestType"`
	Page      int    `json:"Page" form:"Page"`
	Offset    int    `json:"Offset" form:"Offset"`
}
