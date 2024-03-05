package test

import (
	"csw-golang/internal/domain/entity/datastruct"
	"gorm.io/gorm"
)

type ExerciseTestRepo interface {
	AddExerciseTest(exerciseTest datastruct.SubjectTestTypeExercises) (datastruct.SubjectTestTypeExercises, error)
	GetExerciseTest(id string) (datastruct.SubjectTestTypeExercises, error)
	GetAllExerciseTest() ([]datastruct.SubjectTestTypeExercises, error)
}

type exerciseTestRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ExerciseTestRepo {
	return &exerciseTestRepo{
		db,
	}
}
