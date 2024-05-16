package localservice

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"gorm.io/gorm"
)

type QuizService struct {
	DB *gorm.DB
}

func NewQuizService(
	DB *gorm.DB,
) localrepository.Quiz {
	return &QuizService{
		DB,
	}
}

// func (svc *QuizService) FindMentorInfo() ([]*response.MentorQuizList, []int, error) {
// 	mentors := make([]*response.MentorQuizList, 0)

// 	res := svc.DB.
// 		Raw(`SELECT m.uuid, ud.media_id, m.short_name AS name, mdl.name AS teaching_field, m.description, m.motto, COALESCE(ROUND(avg_rating.avg_rating, 2), 0) AS rating
// 			FROM mentors m
// 			LEFT JOIN user_details ud ON ud.user_id = m.user_id
// 			LEFT JOIN modules mdl ON mdl.id = m.module_id
// 			LEFT JOIN (
// 				SELECT mentor_id, AVG(rating) AS avg_rating
// 				FROM user_mentor_testimonials
// 				GROUP BY
// 				mentor_id
// 				) AS avg_rating ON avg_rating.mentor_id = m.id`)

// 	err := res.Scan(&mentors).Error
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to find mentors: %v", err)
// 	}

// 	var mentorMediaIDs []int
// 	for _, mentor := range mentors {
// 		mentorMediaIDs = append(mentorMediaIDs, mentor.MediaID)
// 	}

// 	return mentors, mentorMediaIDs, nil
// }

func (svc *QuizService) FindUserAnswerReview() ([]*response.UserAnswer, error) {
	mentors := make([]*response.UserAnswer, 0)

	res := svc.DB.
		Raw(`select qq.id as question_id,qq.content as question_content,qc.id as choice_id,qc.content as choices_content from quiz_questions as qq inner join quiz_choices as qc on qc.question_id = qq.id where qc.id in(
			select distinct(choice_id) from quiz_answers where submission_id = 2)
		`)

	err := res.Scan(&mentors).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find mentors: %v", err)
	}

	fmt.Println("user answer in locat service")
	for _, mentor := range mentors {
		fmt.Printf("Question ID: %d, Question Content: %s\n", mentor.QuestionId, mentor.QuestionContent)
		fmt.Printf("Choice ID: %d, Choice Content: %s\n", mentor.ChoiceId, mentor.ChoiceContent)
	}
	return mentors, nil
}
