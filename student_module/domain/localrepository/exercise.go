package localrepository

import "github.com/Marvit-Solutions/csw-golang/library/struct/model"

type Exercise interface {
	FindSubModulesID() ([]int, error)
	FindHighestOrLowestScoreAndSubModules(exerciseID int, userID int, findMax bool) (*model.ExerciseSubmissionsModule, []model.ExerciseSubmissionsSubModule, error)
	GetMaxScoreAndTotalQuestionExercise(exerciseID, subModuleID int) (int, int, error)
}
