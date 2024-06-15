package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

type QuizAdmin interface {
	GetQuizAdminAll(searchKeywords string, testTypeID int, subModulID int, limit int, offset int) ([]*response.QuizAdminAllResponse, error)
	CountQuizAdminALL(searchKeywords string, testTypeID int, subModulID int) (int, error)
}
