package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateQuestionQuizzes() []*ds.QuestionQuizzes {
	questionQuizzes := []*ds.QuestionQuizzes{
		// Pancasila
		{
			ID:             "f8259b82-41ca-4d58-90e7-c2be3f4c2352",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "f8259b82-41ca-4d58-90e7-c2be3f4c235f",
			Content:        "",
		},
	}

	return questionQuizzes
}
