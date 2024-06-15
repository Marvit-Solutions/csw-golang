package quiz

import (
	"fmt"
	"math"

	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

func (u *usecase) QuizAll(req request.ParamQuizAll) (*response.QuizAllResponse, int, int, error) {
	limit := req.Limit
	offset := 0

	if req.Page > 0 {
		offset = (req.Page - 1) * limit
	}

	// get the pretest that has been done
	subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
		"uuid": req.SubModuleUUID,
	})

	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find subModule: %v", err)
	}

	module, err := u.modulRepo.FindOneBy(map[string]interface{}{
		"id": subModule.ModuleID,
	})

	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find module: %v", err)
	}

	countQuizAll, err := u.quizLocalRepo.CountQuizALL(req.AuthenticatedUser, req.TestTypeId, subModule.ID)
	totalRows := countQuizAll
	totalPages := int(math.Ceil(float64(countQuizAll) / float64(req.Limit)))
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to countQuizAll: %v", err)
	}

	quizAll, err := u.quizLocalRepo.GetQuizAll(req.AuthenticatedUser, req.TestTypeId, subModule.ID, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find quizAll: %v", err)
	}

	result := &response.QuizAllResponse{
		UUID:          req.SubModuleUUID,
		SubModuleName: subModule.Name,
		ModuleName:    module.Name,
		Quizzes:       quizAll,
	}

	return result, totalRows, totalPages, nil

}
