package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubjectTestTypeQuizzes() []*ds.SubjectTestTypeQuizzes {
	subjectTestTypeQuizzes := []*ds.SubjectTestTypeQuizzes{
		// Pancasila
		{
			ID:        "f8259b82-41ca-4d58-90e7-c2be3f4c235f",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			TestType:  "Module Test",
		},
		{
			ID:        "f8259b82-41ca-4d58-90e7-c2be3f4c235q",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "b8aa7572-a9b6-40e4-9ef5-06cf13f79e49",
			TestType:  "Module Test",
		},
		{
			ID:        "94905d6b-fab1-43f5-ad1e-573b1feea26s",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "94905d6b-fab1-43f5-ad1e-573b1feea26d",
			TestType:  "Module Test",
		},
		{
			ID:        "58cdaedb-c3a0-47b6-96c3-bbb3bad537c7",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58cdaedb-c3a0-47b6-96c3-bbb3bad537c7",
			TestType:  "Module Test",
		},
	}

	return subjectTestTypeQuizzes
}
