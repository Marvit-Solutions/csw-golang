package request

type PretestSubmitRequest struct {
	UserID        string `json:"UserID" form:"UserID"`
	ChoiceQuizzes []struct {
		QuestionQuizID string `json:"QuestionQuizID" form:"QuestionQuizID"`
		ChoiceQuizID   string `json:"ChoiceQuizID" form:"ChoiceQuizID"`
	}
}
type PretestParamRequest struct {
	UserID    string `json:"UserID" form:"UserID"`
	Module    string `json:"Module" form:"Module"`
	SubModule string `json:"SubModule" form:"SubModule"`
	Page      int    `json:"Page" form:"Page"`
	Offset    int    `json:"Offset" form:"Offset"`
}
