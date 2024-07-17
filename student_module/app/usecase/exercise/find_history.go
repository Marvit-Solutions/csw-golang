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
		UUID:                   submissionModuleMax.UUID,
		ScorePerModule:         submissionModuleMax.Score,
		MaxScorePerModule:      0,
		RightAnswerPerModule:   submissionModuleMax.RightAnswer,
		TotalQuestionPerModule: 0,
	}

	perfectScoreModuleMax := 0
	totalQuestionMax := 0
	for _, submissionSubModule := range submissionSubModulesMax {
		// Mendapatkan total max score/perfecet score dan row count/total question exercise_question per sub module
		maxScore, totalQuestion, err := u.exerciseLocalRepo.GetMaxScoreAndTotalQuestionExercise(submissionSubModule.ExerciseID, submissionSubModule.SubModuleID)
		perfectScoreModuleMax += maxScore
		totalQuestionMax += totalQuestion
		if err != nil {
			return nil, fmt.Errorf("failed inf GetMaxScoreAndTotalQuestionExercise: %v", err)
		}
		subModuleDetail := &response.ExerciseHistoryDetailSubModule{
			SubmissionUUID:               submissionModuleMax.UUID,
			SubModule:                    subModuleMap[submissionSubModule.SubModuleID],
			ScorePerSubModule:            submissionSubModule.Score,
			MaxScorePerSubModule:         maxScore,
			TotalRightAnswerPerSubModule: submissionSubModule.RightAnswer,
			TotalQuestionPerSubModule:    totalQuestion,
		}
		historyDetailMax.ExerciseHistoryDetail = append(historyDetailMax.ExerciseHistoryDetail, subModuleDetail)
	}

	historyDetailMax.MaxScorePerModule = perfectScoreModuleMax
	historyDetailMax.TotalQuestionPerModule = totalQuestionMax

	//---------------------------------------- Mencari nilai terendah
	submissionModuleMin, submissionSubModulesMin, err := u.exerciseLocalRepo.FindHighestOrLowestScoreAndSubModules(exercise.ID, false)
	if err != nil {
		fmt.Println("Error finding lowest score submissionModule:", err)
		return nil, fmt.Errorf("failed to find exercise submissionModules: %v", err)

	}
	fmt.Printf("Lowest score submissionModule: %+v\n", submissionModuleMin)
	fmt.Printf("Sub modules: %+v\n", submissionSubModulesMin)

	historyDetailMin := &response.ExerciseHistoryDetailModule{
		UUID:                   submissionModuleMin.UUID,
		ScorePerModule:         submissionModuleMin.Score,
		MaxScorePerModule:      0,
		RightAnswerPerModule:   submissionModuleMin.RightAnswer,
		TotalQuestionPerModule: 0,
	}

	perfectScoreModuleMin := 0
	totalQuestionMin := 0

	for _, submissionSubModule := range submissionSubModulesMin {
		// Mendapatkan total max score dan row count exercise_question per sub module
		maxScore, totalQuestion, err := u.exerciseLocalRepo.GetMaxScoreAndTotalQuestionExercise(submissionSubModule.ExerciseID, submissionSubModule.SubModuleID)
		if err != nil {
			return nil, fmt.Errorf("failed inf GetMaxScoreAndTotalQuestionExercise: %v", err)
		}
		perfectScoreModuleMin += maxScore
		totalQuestionMin += totalQuestion

		subModuleDetail := &response.ExerciseHistoryDetailSubModule{
			SubmissionUUID:               submissionModuleMax.UUID,
			SubModule:                    subModuleMap[submissionSubModule.SubModuleID],
			ScorePerSubModule:            submissionSubModule.Score,
			MaxScorePerSubModule:         maxScore,
			TotalRightAnswerPerSubModule: submissionSubModule.RightAnswer,
			TotalQuestionPerSubModule:    totalQuestion,
		}
		historyDetailMin.ExerciseHistoryDetail = append(historyDetailMin.ExerciseHistoryDetail, subModuleDetail)
	}

	historyDetailMin.MaxScorePerModule = perfectScoreModuleMin
	historyDetailMin.TotalQuestionPerModule = totalQuestionMin

	//-------------------------------------------------------------------------------------
	res := &response.ExerciseHistory{
		MaxScore: historyDetailMax,
		MinScore: historyDetailMin,
	}

	return res, nil
}
