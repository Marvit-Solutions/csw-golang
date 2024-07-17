package localservice

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localrepository"
	"gorm.io/gorm"
)

type ExerciseService struct {
	DB *gorm.DB
}

func NewExerciseService(
	DB *gorm.DB,
) localrepository.Exercise {
	return &ExerciseService{
		DB,
	}
}

func (svc *ExerciseService) FindSubModulesID() ([]int, error) {
	query := `SELECT ARRAY_AGG(DISTINCT sub_module_id ORDER BY sub_module_id) FROM exercise_questions`
	var subModuleIDsStr string

	res := svc.DB.Raw(query).Scan(&subModuleIDsStr)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find user answers: %v", res.Error)
	}

	subModuleIDsStr = strings.Trim(subModuleIDsStr, "{}")
	subModuleIDStrings := strings.Split(subModuleIDsStr, ",")
	subModuleIDs := make([]int, len(subModuleIDStrings))

	for i, idStr := range subModuleIDStrings {
		var id int
		_, err := fmt.Sscanf(idStr, "%d", &id)
		if err != nil {
			return nil, fmt.Errorf("failed to parse subModuleID %s: %v", idStr, err)
		}
		subModuleIDs[i] = id
	}

	for _, subModuleID := range subModuleIDs {
		fmt.Printf("subModuleID: %d\n", subModuleID)
	}
	return subModuleIDs, nil
}

// func (svc *ExerciseService) FindHighestOrLowestScoreAndSubModules(exerciseID int, findMax bool) (*model.ExerciseSubmissionsModule, []model.ExerciseSubmissionsSubModule, error) {
// 	var submission model.ExerciseSubmissionsModule
// 	order := "score DESC"
// 	if !findMax {
// 		order = "score ASC"
// 	}
// 	res := svc.DB.Where("exercise_id = ?", exerciseID).Order(order).First(&submission)
// 	if res.Error != nil {
// 		scoreType := "highest"
// 		if !findMax {
// 			scoreType = "lowest"
// 		}
// 		return nil, nil, fmt.Errorf("failed to find %s score submission: %v", scoreType, res.Error)
// 	}

// 	var subModules []model.ExerciseSubmissionsSubModule
// 	res = svc.DB.Where("submissions_module_id = ?", submission.ID).Find(&subModules)
// 	if res.Error != nil {

// 		return nil, nil, fmt.Errorf("failed to find sub modules: %v", res.Error)
// 	}

// 	return &submission, subModules, nil
// }

func (svc *ExerciseService) FindHighestOrLowestScoreAndSubModules(exerciseID int, findMax bool) (*model.ExerciseSubmissionsModule, []model.ExerciseSubmissionsSubModule, error) {
	var submission model.ExerciseSubmissionsModule
	order := "score DESC"
	if !findMax {
		order = "score ASC"
	}
	res := svc.DB.Where("exercise_id = ?", exerciseID).Order(order).First(&submission)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// If no record is found, return empty data instead of an error
			return &model.ExerciseSubmissionsModule{}, []model.ExerciseSubmissionsSubModule{}, nil
		}
		scoreType := "highest"
		if !findMax {
			scoreType = "lowest"
		}
		return nil, nil, fmt.Errorf("failed to find %s score submission: %v", scoreType, res.Error)
	}

	var subModules []model.ExerciseSubmissionsSubModule
	res = svc.DB.Where("submissions_module_id = ?", submission.ID).Find(&subModules)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// If no submodules are found, return empty data instead of an error
			return &submission, []model.ExerciseSubmissionsSubModule{}, nil
		}
		return nil, nil, fmt.Errorf("failed to find sub modules: %v", res.Error)
	}

	return &submission, subModules, nil
}

func (svc *ExerciseService) GetMaxScoreAndTotalQuestionExercise(exerciseID, subModuleID int) (int, int, error) {
	var totalScore int
	var rowCount int64

	// Menghitung total score
	res := svc.DB.Model(&model.ExerciseQuestion{}).
		Where("exercise_id = ? AND sub_module_id = ?", exerciseID, subModuleID).
		Select("COALESCE(SUM(score), 0)").Scan(&totalScore)
	if res.Error != nil {
		return 0, 0, fmt.Errorf("failed to calculate total score: %v", res.Error)
	}

	// Menghitung total row
	res = svc.DB.Model(&model.ExerciseQuestion{}).
		Where("exercise_id = ? AND sub_module_id = ?", exerciseID, subModuleID).
		Count(&rowCount)
	if res.Error != nil {
		return 0, 0, fmt.Errorf("failed to count rows: %v", res.Error)
	}

	return totalScore, int(rowCount), nil
}
