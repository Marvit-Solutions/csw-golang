package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"gorm.io/gorm"
	"time"
)

func CreateExerciseType() []*ds.SubjectTestTypeExercises {
	return []*ds.SubjectTestTypeExercises{{
		ID:        "9975116b-01a4-4978-90aa-a3a7525d504a",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		TestType:  "try-out1",
	},
	}
}
