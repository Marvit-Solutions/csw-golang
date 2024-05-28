package quiz

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

func (u *usecase) QuizSubModuleAll(req request.ParamQuizSubModuleAll) (*response.QuizSubModuleAllResponse, error) {

	modul, err := u.modulRepo.FindOneBy(map[string]interface{}{
		"id": req.ModuleID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find modul: %v", err)
	}

	QuizzesGroupedBySubModule, err := u.quizLocalRepo.CountQuizzesGroupedBySubModule(modul.ID, req.TestTypeId)

	if err != nil {
		return nil, fmt.Errorf("failed to find QuizzesGroupedBySubModule: %v", err)
	}

	result := &response.QuizSubModuleAllResponse{
		UUID:                      modul.UUID,
		ModuleName:                modul.Name,
		QuizzesGroupedBySubModule: QuizzesGroupedBySubModule,
	}

	return result, nil
}
