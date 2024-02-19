package module

import (
	"csw-golang/internal/domain/entity/datastruct"

	"gorm.io/gorm"
)

type ModuleRepo interface {
	GetListModules() ([]datastruct.SubModules, error)
	GetListSubjectQuiz() ([]datastruct.SubjectTestTypeQuizzes, error)
	GetSubjectsBySubmoduleID(submoduleID string) ([]datastruct.Subjects, error)
	GetTestByTestTypeID(testTypeID string) (datastruct.SubjectTestTypeQuizzes, error)
	GetTestReview(moduleTestId string) (datastruct.SubjectTestTypeQuizzes, error)
	GetSubmittedReviewByID(moduleTestId string) (datastruct.UserTestSubmissionQuizzes, error)
	AddQuizSubmission(submission datastruct.UserTestSubmissionQuizzes) (string, error)
	PostSubmittedQuizAnswer(testTypeID string, submittedAnswers []datastruct.UserSubmittedAnswerQuizzes) error
	AddGrade(grade datastruct.GradeQuizzes) error
	GetTop3EverySubject(userID string, subjectTypeID string) ([]datastruct.SubModules, error)
}

type moduleRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ModuleRepo {
	return &moduleRepo{
		db,
	}
}
