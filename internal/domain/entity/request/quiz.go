package request

type QuizParamRequest struct {
	UserID    string `json:"UserID" form:"UserID"`
	Module    string `json:"Module" form:"Module"`
	SubModule string `json:"SubModule" form:"SubModule"`
	TestType  string `json:"TestType" form:"TestType"`
	Limit     int64  `json:"Limit" form:"Limit"`
	Offset    int    `json:"Offset" form:"Offset"`
}
