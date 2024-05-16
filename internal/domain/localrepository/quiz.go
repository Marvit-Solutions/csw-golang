package localrepository

import "github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"

type Quiz interface {
	FindUserAnswerReview() ([]*response.UserAnswer, error)
}
