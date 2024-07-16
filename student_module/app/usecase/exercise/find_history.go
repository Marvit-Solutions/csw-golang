package exercise

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

func (u *usecase) FindHistory(req request.ExerciseHistory) (*response.ExerciseHistory, error) {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id":      req.AuthenticatedUser,
		"role_id": helper.PembeliPaketBimbel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	if user == nil {
		return nil, helper.ErrAccessDenied
	}

	exercise, err := u.exerciseRepo.FindOneBy(map[string]interface{}{
		"uuid": req.ExerciseUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise: %v", err)
	}

	subModules, err := u.subModuleRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find sub modules: %v", err)
	}

	subModuleMap := make(map[int]string)
	for _, subModule := range subModules {
		subModuleMap[subModule.ID] = subModule.Name
	}
	// exerciseSubmissions, err := u.exerciseSubmissionsSubModuleRepo.FindBy(map[string]interface{}{
	// 	"exercise_id": exercise.ID,
	// 	"user_id":     req.AuthenticatedUser,
	// }, 0, 0)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to find exercise submissions: %v", err)
	// }

	//---------------------------------------- Mencari nilai tertinggi
	submissionModuleMax, submissionSubModulesMax, err := u.exerciseLocalRepo.FindHighestOrLowestScoreAndSubModules(exercise.ID, true)
	if err != nil {
		fmt.Println("Error finding highest score submissionModule:", err)
		return nil, fmt.Errorf("failed to find exercise submissionModules: %v", err)

	}
	fmt.Printf("Highest score submissionModule: %+v\n", submissionModuleMax)
	fmt.Printf("Sub modules: %+v\n", submissionSubModulesMax)

	historyDetailMax := &response.ExerciseHistoryDetailModule{
		ScorePerModule:    submissionModuleMax.Score,
		MaxScorePerModule: submissionModuleMax.RightAnswer, // As an example, assuming RightAnswer represents the max score
	}

	for _, submissionSubModule := range submissionSubModulesMax {
		// Mendapatkan total max score dan row count exercise_question per sub module
		MaxScore, TotalQuestion, err := u.exerciseLocalRepo.GetMaxScoreAndTotalQuestionExercise(submissionSubModule.ExerciseID, submissionSubModule.SubModuleID)
		if err != nil {
			return nil, fmt.Errorf("failed inf GetMaxScoreAndTotalQuestionExercise: %v", err)
		}
		subModuleDetail := &response.ExerciseHistoryDetailSubModule{
			SubmissionUUID:               submissionModuleMax.UUID,
			SubModule:                    subModuleMap[submissionSubModule.SubModuleID],
			ScorePerSubModule:            submissionSubModule.Score,
			MaxScorePerSubModule:         MaxScore,
			TotalRightAnswerPerSubModule: submissionSubModule.RightAnswer,
			MaxTotalQuestionPerSubModule: TotalQuestion,
		}
		historyDetailMax.ExerciseHistoryDetail = append(historyDetailMax.ExerciseHistoryDetail, subModuleDetail)
	}

	//---------------------------------------- Mencari nilai terendah
	submissionModuleMin, submissionSubModulesMin, err := u.exerciseLocalRepo.FindHighestOrLowestScoreAndSubModules(exercise.ID, false)
	if err != nil {
		fmt.Println("Error finding lowest score submissionModule:", err)
		return nil, fmt.Errorf("failed to find exercise submissionModules: %v", err)

	}
	fmt.Printf("Lowest score submissionModule: %+v\n", submissionModuleMin)
	fmt.Printf("Sub modules: %+v\n", submissionSubModulesMin)

	historyDetailMin := &response.ExerciseHistoryDetailModule{
		ScorePerModule:    submissionModuleMin.Score,
		MaxScorePerModule: submissionModuleMin.RightAnswer,
	}

	for _, submissionSubModule := range submissionSubModulesMin {
		// Mendapatkan total max score dan row count exercise_question per sub module
		MaxScore, TotalQuestion, err := u.exerciseLocalRepo.GetMaxScoreAndTotalQuestionExercise(submissionSubModule.ExerciseID, submissionSubModule.SubModuleID)
		if err != nil {
			return nil, fmt.Errorf("failed inf GetMaxScoreAndTotalQuestionExercise: %v", err)
		}
		subModuleDetail := &response.ExerciseHistoryDetailSubModule{
			SubmissionUUID:               submissionModuleMax.UUID,
			SubModule:                    subModuleMap[submissionSubModule.SubModuleID],
			ScorePerSubModule:            submissionSubModule.Score,
			MaxScorePerSubModule:         MaxScore,
			TotalRightAnswerPerSubModule: submissionSubModule.RightAnswer,
			MaxTotalQuestionPerSubModule: TotalQuestion,
		}
		historyDetailMin.ExerciseHistoryDetail = append(historyDetailMin.ExerciseHistoryDetail, subModuleDetail)
	}

	//-------------------------------------------------------------------------------------
	res := &response.ExerciseHistory{
		MaxScore: historyDetailMax,
		MinScore: historyDetailMin,
	}

	return res, nil
}
