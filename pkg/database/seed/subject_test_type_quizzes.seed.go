package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubjectTestTypeQuizzes() []*ds.SubjectTestTypeQuizzes {
	subjectTestTypeQuizzes := []*ds.SubjectTestTypeQuizzes{
		{
			ID:          "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubjectID:   "58ab4e13-e016-4c42-8123-74b220db465a",
			TestType:    "PRETEST",
			Title:       "Pretest Pancasila",
			MeetingDate: time.Now().AddDate(0, 0, 1),
			Open:        time.Now(),
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book...",
			Time:        15,
			Point:       100,
			Attempt:     1,
		},
		{
			ID:          "8049dc26-01c1-426a-803d-570b799d3810",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubjectID:   "58ab4e13-e016-4c42-8123-74b220db465a",
			TestType:    "POSTTEST ",
			Title:       "Posttest Pancasila",
			MeetingDate: time.Now().AddDate(0, 0, 1),
			Open:        time.Now(),
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book...",
			Time:        15,
			Point:       100,
			Attempt:     1,
		},
		{
			ID:          "b1d5958d-96fc-4b85-8f83-78e4eb8ffbbc",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubjectID:   "b8aa7572-a9b6-40e4-9ef5-06cf13f79e49",
			TestType:    "PRETEST",
			Title:       "Pretest Konstitusi yang berlaku di Indonesia",
			MeetingDate: time.Now().AddDate(0, 0, 1),
			Open:        time.Now(),
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book...",
			Time:        15,
			Point:       100,
			Attempt:     1,
		},
		{
			ID:          "2b703858-1fbd-49ad-8574-4ff5962dc92c",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubjectID:   "b8aa7572-a9b6-40e4-9ef5-06cf13f79e49",
			TestType:    "POSTTEST ",
			Title:       "Posttest Konstitusi yang berlaku di Indonesia",
			MeetingDate: time.Now().AddDate(0, 0, 1),
			Open:        time.Now(),
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book...",
			Time:        15,
			Point:       100,
			Attempt:     1,
		},
		{
			ID:          "68acf1ca-55d2-4666-9f7d-4ed32eac62d8",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubjectID:   "94905d6b-fab1-43f5-ad1e-573b1feea26d",
			TestType:    "PRETEST",
			Title:       "Pretest NKRI & Bhinneka Tunggal Ika",
			MeetingDate: time.Now().AddDate(0, 0, 1),
			Open:        time.Now(),
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book...",
			Time:        15,
			Point:       100,
			Attempt:     1,
		},
		{
			ID:          "aba04800-d0c7-46de-ba87-32809d651d02",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubjectID:   "94905d6b-fab1-43f5-ad1e-573b1feea26d",
			TestType:    "POSTTEST ",
			Title:       "Posttest NKRI & Bhinneka Tunggal Ika",
			MeetingDate: time.Now().AddDate(0, 0, 1),
			Open:        time.Now(),
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book...",
			Time:        15,
			Point:       100,
			Attempt:     1,
		},
	}

	return subjectTestTypeQuizzes
}
