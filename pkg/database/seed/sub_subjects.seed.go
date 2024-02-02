package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubSubjects() []*ds.SubSubject {
	subSubjects := []*ds.SubSubject{
		// Pancasila
		{
			ID:        "f8259b82-41ca-4d58-90e7-c2be3f4c235f",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			Name:      "Pengertian Pancasila",
			Content:   "When an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		},
		{
			ID:        "69c1abba-6832-4857-9723-92c3f8efa1bc",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			Name:      "Asal Usul Pancasila",
			Content:   "When an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		},
		{
			ID:        "3540e42b-e834-46a9-b2d5-555a820ddbce",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			Name:      "Sejarah Pancasila",
			Content:   "When an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		},
		{
			ID:        "3fe04671-ef4b-4d1a-bd61-29aeff6bece9",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			Name:      "Lambang dan Nilai Pancasila",
			Content:   "When an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		},
		{
			ID:        "bed41b17-58f2-4ccb-91d3-4ebb275f39fc",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			Name:      "Dimensi Pancasila",
			Content:   "When an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		},
		{
			ID:        "2fab6d76-13b9-4454-b66f-8925371fb32e",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			SubjectID: "58ab4e13-e016-4c42-8123-74b220db465a",
			Name:      "Rumusan Kesatuan Sila Pancasila",
			Content:   "When an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		},
	}

	return subSubjects
}
