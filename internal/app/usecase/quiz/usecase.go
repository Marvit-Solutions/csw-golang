package quiz

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localservice"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	QuizContent(req request.ParamQuizContent) (*response.QuizContentResponse, error)
	QuizReview(req request.ParamQuizReview) (*response.QuizReviewResponse, error)
	QuizSubmission(req request.QuizSubmissionRequest) error
	QuizDetail(req request.ParamQuizDetail) (*response.QuizDetailResponse, error)
	QuizScoreAll(req request.ParamQuizScoreAll) (*response.QuizScoreAllResponse, error)
	QuizSubModuleAll(req request.ParamQuizSubModuleAll) (*response.QuizSubModuleAllResponse, error)
	QuizAll(req request.ParamQuizAll) (*response.QuizAllResponse, int, int, error)
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
	quizLocalRepo      localrepository.Quiz
	subModulRepo       repository.SubModuleRepository
	modulRepo          repository.ModuleRepository
	userRepo           repository.UserRepository
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
		quizLocalRepo:      localservice.NewQuizService(db),
		subModulRepo:       service.NewSubModuleService(db, nil),
		modulRepo:          service.NewModuleService(db, nil),
		userRepo:           service.NewUserService(db, nil),
	}
}
