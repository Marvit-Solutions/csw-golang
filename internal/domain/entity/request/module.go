package request

type UserSubmittedQuizRequest struct {
	// UserID           string             `json:"UserID" form:"UserID"`
	PairOfUserAnswer []PairOfUserAnswer `json:"PairOfUserAnswer"`
}

type PairOfUserAnswer struct {
	QuestionQuizID string `json:"QuestionQuizID" form:"QuestionQuizID"`
	ChoiceQuizID   string `json:"ChoiceQuizID" form:"ChoiceQuizID"`
}
