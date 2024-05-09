package quiz

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	QuizContent(req request.ParamQuiz) (*response.QuizContentResponse, error)
	QuizSubmission(req request.QuizSubmissionRequest) error
	DetailQuiz(req request.ParamDetailQuiz) (*response.DetailQuizResponse, error)
}

type usecase struct {
	db                 *gorm.DB
	quizRepo           repository.QuizRepository
	subjectRepo        repository.SubjectRepository
	testTypeRepo       repository.TestTypeRepository
	quizQuestionRepo   repository.QuizQuestionRepository
	quizChoiceRepo     repository.QuizChoiceRepository
	quizAnswerRepo     repository.QuizAnswerRepository
	quizSubmissionRepo repository.QuizSubmissionRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                 db,
		quizRepo:           service.NewQuizService(db, nil),
		subjectRepo:        service.NewSubjectService(db, nil),
		testTypeRepo:       service.NewTestTypeService(db, nil),
		quizQuestionRepo:   service.NewQuizQuestionService(db, nil),
		quizChoiceRepo:     service.NewQuizChoiceService(db, nil),
		quizAnswerRepo:     service.NewQuizAnswerService(db, nil),
		quizSubmissionRepo: service.NewQuizSubmissionService(db, nil),
	}
}
