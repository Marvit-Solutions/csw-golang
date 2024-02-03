package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubjectTestTypeQuizzes() []*ds.SubjectTestTypeQuizzes {
	subjectTestTypeQuizzes := []*ds.SubjectTestTypeQuizzes{
		{
			ID:        "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			TestType:  "Pretest",
		},
		{
			ID:        "8049dc26-01c1-426a-803d-570b799d3810",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			TestType:  "Posttest ",
		},
		{
			ID:        "b1d5958d-96fc-4b85-8f83-78e4eb8ffbbc",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "b8aa7572-a9b6-40e4-9ef5-06cf13f79e49",
			TestType:  "Pretest",
		},
		{
			ID:        "2b703858-1fbd-49ad-8574-4ff5962dc92c",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "b8aa7572-a9b6-40e4-9ef5-06cf13f79e49",
			TestType:  "Posttest ",
		},
		{
			ID:        "68acf1ca-55d2-4666-9f7d-4ed32eac62d8",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "94905d6b-fab1-43f5-ad1e-573b1feea26d",
			TestType:  "Pretest",
		},
		{
			ID:        "aba04800-d0c7-46de-ba87-32809d651d02",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "94905d6b-fab1-43f5-ad1e-573b1feea26d",
			TestType:  "Posttest ",
		},
	}

	return subjectTestTypeQuizzes
}
