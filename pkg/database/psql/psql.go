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
		&ds.UserDetail{},
		&ds.User{},
		&ds.Role{},
		&ds.Address{},
		&ds.Plan{},
		&ds.SubPlan{},
		&ds.SubPlanDetail{},
		&ds.Module{},
		&ds.SubModule{},
		&ds.Subject{},
		&ds.SubSubject{},
		&ds.SubjectTestType{},
		&ds.Question{},
		&ds.Choice{},
		&ds.UserTestSubmission{},
		&ds.UserSubmittedAnswer{},
		&ds.Grade{},
	// 	// &ds.Transaction{},
	// 	// &ds.Subscription{},
	// 	// &ds.Mentor{},
	// 	// &ds.Testimonials{},
	)
	DB.AutoMigrate(
		ds.Role{},
		ds.User{},
		ds.UserDetail{},
		ds.Address{},
		ds.Plan{},
		ds.SubPlan{},
		ds.SubPlanDetail{},
		ds.Module{},
		ds.SubModule{},
		ds.Subject{},
		ds.SubSubject{},
		ds.SubjectTestType{},
		ds.Question{},
		ds.Choice{},
		ds.UserTestSubmission{},
		ds.UserSubmittedAnswer{},
		ds.Grade{},
		// ds.Transaction{},
		// ds.Subscription{},
		// ds.Mentor{},
		// ds.Testimonials{},
	)
	DB.Migrator().HasConstraint(&ds.User{}, "UserDetail")
	DB.Migrator().HasConstraint(&ds.Role{}, "Users")
	DB.Migrator().HasConstraint(&ds.UserDetail{}, "Alamat")
	DB.Migrator().HasConstraint(&ds.Plan{}, "SubPlan")
	DB.Migrator().HasConstraint(&ds.SubPlan{}, "SubPlanDetail")
	DB.Migrator().HasConstraint(&ds.Module{}, "Module")
	DB.Migrator().HasConstraint(&ds.SubModule{}, "SubModule")
	DB.Migrator().HasConstraint(&ds.Subject{}, "Subject")
	DB.Migrator().HasConstraint(&ds.SubSubject{}, "SubSubject")
	DB.Migrator().HasConstraint(&ds.SubjectTestType{}, "SubjectTestType")
	DB.Migrator().HasConstraint(&ds.Question{}, "Question")
	DB.Migrator().HasConstraint(&ds.Choice{}, "Choice")
	DB.Migrator().HasConstraint(&ds.UserTestSubmission{}, "UserTestSubmission")
	DB.Migrator().HasConstraint(&ds.UserSubmittedAnswer{}, "UserSubmittedAnswer")
	DB.Migrator().HasConstraint(&ds.Grade{}, "Grade")

	// DB.Migrator().HasConstraint(&ds.Transaction{}, "Transaksi")
	// DB.Migrator().HasConstraint(&ds.Subscription{}, "Subskripsi")
	// DB.Migrator().HasConstraint(&ds.Mentor{}, "mentor")
	// DB.Migrator().HasConstraint(&ds.Mentor{}, "testimonials")
}
