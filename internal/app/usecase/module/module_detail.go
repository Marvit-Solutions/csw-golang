package module

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
)

func (u *usecase) ModuleDetail(req request.ParamModule) (*response.ModuleDetailResponse, error) {
	subModule, err := u.subModuleRepo.FindOneBy(map[string]interface{}{
		"uuid": req.SubModuleUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find sub module: %v", err)
	}

	module, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"id": subModule.ModuleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	subjects, err := u.subjectRepo.FindBy(map[string]interface{}{
		"sub_module_id": subModule.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find subjects: %v", err)
	}

	subjectIDs := make([]int, len(subjects))
	for i, subject := range subjects {
		subjectIDs[i] = subject.ID
	}

	testType, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"slug": "latihan-soal",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find test type: %v", err)
	}

	quizzes, err := u.quizRepo.FindBy(map[string]interface{}{
		"subject_id":   subjectIDs,
		"test_type_id": &testType.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find quizes: %v", err)
	}

	result := &response.ModuleDetailResponse{
		UUID:          subModule.UUID,
		SubModuleName: subModule.Name,
		Description:   subModule.Description,
		ModuleName:    module.Name,
		Subjects:      make([]*response.SubjectResponse, len(subjects)),
		Quizzes:       make([]*response.QuizResponse, len(quizzes)),
	}

	for i, subject := range subjects {
		result.Subjects[i] = &response.SubjectResponse{
			ID:   subject.ID,
			UUID: subject.UUID,
			Name: subject.Name,
		}
	}

	for i, quiz := range quizzes {
		result.Quizzes[i] = &response.QuizResponse{
			ID:    quiz.ID,
			UUID:  quiz.UUID,
			Title: quiz.Title,
		}
	}

	sort.Slice(result.Subjects, func(i, j int) bool {
		return result.Subjects[i].ID < result.Subjects[j].ID
	})

	sort.Slice(result.Quizzes, func(i, j int) bool {
		return result.Quizzes[i].ID < result.Quizzes[j].ID
	})

	return result, nil
}
