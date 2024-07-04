package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

type Dashboard interface {
	GetQuizAllDashboard(authenticatedUserID int, testTypeID int, modulID int) ([]*response.QuizItemAllDashboard, error)
}
