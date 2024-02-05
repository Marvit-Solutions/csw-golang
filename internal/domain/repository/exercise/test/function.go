package test

import (
	"csw-golang/internal/domain/entity/datastruct"
)

func (e exerciseTestRepo) AddExerciseTest(exerciseTest datastruct.SubjectTestTypeExercises) (datastruct.SubjectTestTypeExercises, error) {

	tx := e.db.Create(&exerciseTest)

	if tx.Error != nil {
		return datastruct.SubjectTestTypeExercises{}, tx.Error
	}
	return exerciseTest, nil
}

func (e exerciseTestRepo) GetExerciseTest(id string) (datastruct.SubjectTestTypeExercises, error) {
	var exercises datastruct.SubjectTestTypeExercises

	tx := e.db.Preload("QuestionExercises.ChoiceExercises").Where("id = ?", id).First(&exercises)
	if tx.Error != nil {
		return datastruct.SubjectTestTypeExercises{}, tx.Error
	}
	return exercises, nil
}

func (e exerciseTestRepo) GetAllExerciseTest() ([]datastruct.SubjectTestTypeExercises, error) {
	var exercises []datastruct.SubjectTestTypeExercises

	tx := e.db.Find(&exercises)
	if tx.Error != nil {
		return []datastruct.SubjectTestTypeExercises{}, tx.Error
	}
	return exercises, nil
}
