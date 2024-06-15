package quizAdmin

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
	QuizAdminAll(req request.ParamQuizAdminAll) ([]*response.QuizAdminAllResponse, int, int, error)
	Create(req request.QuizAdminPayload) error
	Update(req request.QuizAdminPayloadUpdate) error
	Delete(uuid string) error
	QuizAdminUpdateDetail(req request.ParamQuizAdminUpdateDetail) (*response.QuizAdminUpdateDetailResponse, error)
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
	quizAdminLocalRepo localrepository.QuizAdmin
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
		quizAdminLocalRepo: localservice.NewQuizAdminService(db),
		subModulRepo:       service.NewSubModuleService(db, nil),
		modulRepo:          service.NewModuleService(db, nil),
		userRepo:           service.NewUserService(db, nil),
	}
}
