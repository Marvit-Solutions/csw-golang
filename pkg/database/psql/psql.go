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
		ds.SubjectTestType{},
		ds.Questions{},
		ds.Choices{},
		ds.UserTestSubmissions{},
		ds.UserSubmittedAnswers{},
		ds.Grades{},
	// 	// &ds.Transaction{},
	// 	// &ds.Subscription{},
	// 	// &ds.Mentor{},
	// 	// &ds.Testimonials{},
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
		ds.SubjectTestType{},
		ds.Questions{},
		ds.Choices{},
		ds.UserTestSubmissions{},
		ds.UserSubmittedAnswers{},
		ds.Grades{},
		// ds.Transaction{},
		// ds.Subscription{},
		// ds.Mentor{},
		// ds.Testimonials{},
	)
	DB.Migrator().HasConstraint(&ds.Users{}, "UserDetail")
	DB.Migrator().HasConstraint(&ds.Roles{}, "Users")
	DB.Migrator().HasConstraint(&ds.UserDetails{}, "Alamat")
	DB.Migrator().HasConstraint(&ds.Modules{}, "Module")
	DB.Migrator().HasConstraint(&ds.Modules{}, "SubPlan")
	DB.Migrator().HasConstraint(&ds.SubModules{}, "SubModule")
	DB.Migrator().HasConstraint(&ds.Subjects{}, "Subject")
	DB.Migrator().HasConstraint(&ds.SubSubject{}, "SubSubject")
	DB.Migrator().HasConstraint(&ds.SubjectTestType{}, "SubjectTestType")
	DB.Migrator().HasConstraint(&ds.Questions{}, "Question")
	DB.Migrator().HasConstraint(&ds.Choices{}, "Choice")
	DB.Migrator().HasConstraint(&ds.UserTestSubmissions{}, "UserTestSubmission")
	DB.Migrator().HasConstraint(&ds.UserSubmittedAnswers{}, "UserSubmittedAnswer")
	DB.Migrator().HasConstraint(&ds.Grades{}, "Grade")

	// DB.Migrator().HasConstraint(&ds.Transaction{}, "Transaksi")
	// DB.Migrator().HasConstraint(&ds.Subscription{}, "Subskripsi")
	// DB.Migrator().HasConstraint(&ds.Mentor{}, "mentor")
	// DB.Migrator().HasConstraint(&ds.Mentor{}, "testimonials")
}
