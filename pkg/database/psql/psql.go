package psql

import (
	"fmt"
	"os"

	ds "csw-golang/internal/domain/entity/datastruct"
	"csw-golang/pkg/database/seed"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
	seed.DBSeed(DB)
}

func InitDB() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		username,
		password,
		name,
		host,
		port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.Migrator().DropTable(
		ds.Roles{},
		ds.Users{},
		ds.UserDetails{},
		ds.Addresses{},
		ds.SubPlans{},
		ds.Modules{},
		ds.SubModules{},
		ds.Subjects{},
		ds.SubSubject{},
		ds.SubjectTestTypeQuizzes{},
		//ds.SubjectTestTypeExercises{},
		ds.QuestionQuizzes{},
		//ds.QuestionExercises{},
		//ds.ChoiceQuizzes{},
		//ds.ChoiceExercises{},
		ds.UserTestSubmissionQuizzes{},
		//ds.UserTestSubmissionExercises{},
		ds.UserSubmittedAnswerQuizzes{},
		//ds.UserSubmittedAnswerExercises{},
		//ds.GradeExercises{},
		ds.GradeQuizzes{},
		ds.Testimonials{},
		ds.Mentors{},
		//ds.QuestionExercises{},
		// &ds.Transaction{},
		// &ds.Subscription{},
	)
	DB.AutoMigrate(
		ds.Roles{},
		ds.Users{},
		ds.UserDetails{},
		ds.Addresses{},
		ds.SubPlans{},
		ds.Modules{},
		ds.SubModules{},
		ds.Subjects{},
		ds.SubSubject{},
		ds.SubjectTestTypeQuizzes{},
		ds.SubjectTestTypeExercises{},
		ds.QuestionQuizzes{},
		//ds.QuestionExercises{},
		ds.ChoiceQuizzes{},
		ds.ChoiceExercises{},
		ds.UserTestSubmissionQuizzes{},
		ds.UserTestSubmissionExercises{},
		ds.UserSubmittedAnswerQuizzes{},
		ds.UserSubmittedAnswerExercises{},
		ds.GradeExercises{},
		ds.GradeQuizzes{},
		ds.Testimonials{},
		ds.Mentors{},
		//ds.QuestionExercises{},
		// ds.Transaction{},
		// ds.Subscription{},
	)
}
