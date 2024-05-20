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
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	subModul, err := u.subModulRepo.FindBy(map[string]interface{}{
		"module_id": modul.ID,
	}, 0, 0)

	if err != nil {
		return nil, fmt.Errorf("failed to find sub modules: %v", err)
	}

	fmt.Println(subModul)

	QuizzesGroupedBySubModule, err := u.quizLocalRepo.CountQuizzesGroupedBySubModule(modul.ID, req.TestTypeId)

	if err != nil {
		return nil, fmt.Errorf("failed to select QuizzesGroupedBySubModule: %v", err)
	}

	result := &response.QuizSubModuleAllResponse{
		UUID:                      modul.UUID,
		ModuleName:                modul.Name,
		QuizzesGroupedBySubModule: QuizzesGroupedBySubModule,
	}

	// sort.Slice(result.QuizScoreAllItem, func(i, j int) bool {
	// 	return result.QuizScoreAllItem[i].SubjectID < result.QuizScoreAllItem[j].SubjectID
	// })

	return result, nil
}
