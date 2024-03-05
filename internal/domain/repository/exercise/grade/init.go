package grade

import (
	"csw-golang/internal/domain/entity/datastruct"
	"gorm.io/gorm"
)

type ExerciseGradeRepo interface {
	AddGrade(grade datastruct.GradeExercises) (datastruct.GradeExercises, error)
	GetGrade(userID string, submissionID string) ([]datastruct.GradeExercises, error)
	GetGradeByID(gradeID string) (datastruct.GradeExercises, error)
}

type exercisesGradeRepo struct {
	db *gorm.DB
}

func (e exercisesGradeRepo) GetGradeByID(gradeID string) (datastruct.GradeExercises, error) {
	var grade datastruct.GradeExercises
	tx := e.db.Where("id = ?", gradeID).First(&grade)
	if tx.Error != nil {
		return datastruct.GradeExercises{}, nil
	}
	return grade, nil

}

func (e exercisesGradeRepo) GetGrade(userID string, submissionID string) ([]datastruct.GradeExercises, error) {
	var gradeArr []datastruct.GradeExercises
	tx := e.db.Where("user_id = ? and test_type_exercise_id = ?", userID, submissionID).Find(&gradeArr)
	if tx.Error != nil {
		return []datastruct.GradeExercises{}, nil
	}
	return gradeArr, nil
}

func (e exercisesGradeRepo) AddGrade(grade datastruct.GradeExercises) (datastruct.GradeExercises, error) {
	var submitted []datastruct.UserSubmittedAnswerExercises

	tx := e.db.Where("user_test_submission_exercise_id = ?", grade.UserTestSubmissionExerciseID).Find(&submitted)
	if tx.Error != nil {
		return datastruct.GradeExercises{}, nil
	}

	var score int
	var exercisesArr []datastruct.ChoiceExercises
	for _, ss := range submitted {
		choiceId := ss.ChoiceExerciseID
		var choice datastruct.ChoiceExercises
		_ = e.db.Where("id = ?", choiceId).Find(&choice)
		if choice.IsCorrect == true {
			score += choice.Weight
		}
		exercisesArr = append(exercisesArr, choice)
	}
	grade.Score = score
	create := e.db.Create(&grade)
	if create.Error != nil {
		return datastruct.GradeExercises{}, nil
	}
	return grade, nil
}

func New(db *gorm.DB) ExerciseGradeRepo {
	return &exercisesGradeRepo{
		db,
	}
}
