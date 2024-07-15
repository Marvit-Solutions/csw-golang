package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseSubmissionsSubModule = "exercise_submissions_sub_module"

// ExerciseSubmissionsSubModule mapped from table <exercise_submissions_sub_module>
type ExerciseSubmissionsSubModule struct {
	ID                          int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID                        string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	UserID                      int          `gorm:"column:user_id;not null" json:"user_id"`
	ExerciseID                  int          `gorm:"column:exercise_id;not null" json:"exercise_id"`
	SubModuleID                 int          `gorm:"column:sub_module_id;not null" json:"sub_module_id"`
	ExerciseSubmissionsModuleID int          `gorm:"column:exercise_submissions_module_id;not null" json:"exercise_submissions_module_id"`
	RightAnswer                 int          `gorm:"column:right_answer;not null" json:"right_answer"`
	Score                       int          `gorm:"column:score;not null" json:"score"`
	CreatedAt                   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt                   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt                   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseSubmissionsSubModule's table name
func (*ExerciseSubmissionsSubModule) TableName() string {
	return TableNameExerciseSubmissionsSubModule
}
