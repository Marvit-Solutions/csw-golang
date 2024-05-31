package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

type Dashboard interface {
	GetQuizAllDashboard(authenticatedUserID int, testTypeID int, modulID int) ([]*response.QuizItemAllDashboard, error)
}
