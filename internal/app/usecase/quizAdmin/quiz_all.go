package quizAdmin

import (
	"fmt"
	"math"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

func (u *usecase) QuizAdminAll(req request.ParamQuizAdminAll) ([]*response.QuizAdminAllResponse, int, int, error) {
	limit := req.Limit
	offset := 0

	if req.Page > 0 {
		offset = (req.Page - 1) * limit
	}

	var subModuleID, testTypeID int

	if req.SubModuleName != "" {
		subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
			"name": req.SubModuleName,
		})
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to find subModule: %v", err)
		}
		subModuleID = subModule.ID
	}

	if req.TestTypeName != "" {
		testType, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
			"name": req.TestTypeName,
		})
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to find testType: %v", err)
		}
		testTypeID = testType.ID
	}

	countQuizAll, err := u.quizAdminLocalRepo.CountQuizAdminALL(req.SearchKeywords, testTypeID, subModuleID)
	totalRows := countQuizAll
	totalPages := int(math.Ceil(float64(countQuizAll) / float64(req.Limit)))
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to countQuizAll: %v", err)
	}

	quizAll, err := u.quizAdminLocalRepo.GetQuizAdminAll(req.SearchKeywords, testTypeID, subModuleID, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find quizAll: %v", err)
	}

	result := quizAll

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID > result[j].ID
	})

	return result, totalRows, totalPages, nil
}
