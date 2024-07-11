package subjectAdmin

import (
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localrepository"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localservice"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"gorm.io/gorm"
)

type Usecase interface {
	SubjectAdminAll(req request.ParamSubjectAdminAll) ([]*response.SubjectAdminAllResponse, int, int, error)
	Create(req request.SubjectAdminPayload) error
	Update(req request.SubjectAdminPayloadUpdate) error
	Delete(uuid string) error
	SubjectAdminUpdateDetail(req request.ParamSubjectAdminUpdateDetail) (*response.SubjectAdminUpdateDetailResponse, error)
	Read() ([]*model.Subject, error)
}

type usecase struct {
	db                    *gorm.DB
	quizRepo              repository.QuizRepository
	subjectRepo           repository.SubjectRepository
	testTypeRepo          repository.TestTypeRepository
	quizQuestionRepo      repository.QuizQuestionRepository
	quizChoiceRepo        repository.QuizChoiceRepository
	quizAnswerRepo        repository.QuizAnswerRepository
	quizSubmissionRepo    repository.QuizSubmissionRepository
	subjectAdminLocalRepo localrepository.SubjectAdmin
	subModulRepo          repository.SubModuleRepository
	modulRepo             repository.ModuleRepository
	userRepo              repository.UserRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                    db,
		quizRepo:              service.NewQuizService(db, nil),
		subjectRepo:           service.NewSubjectService(db, nil),
		testTypeRepo:          service.NewTestTypeService(db, nil),
		quizQuestionRepo:      service.NewQuizQuestionService(db, nil),
		quizChoiceRepo:        service.NewQuizChoiceService(db, nil),
		quizAnswerRepo:        service.NewQuizAnswerService(db, nil),
		quizSubmissionRepo:    service.NewQuizSubmissionService(db, nil),
		subjectAdminLocalRepo: localservice.NewSubjectAdminService(db),
		subModulRepo:          service.NewSubModuleService(db, nil),
		modulRepo:             service.NewModuleService(db, nil),
		userRepo:              service.NewUserService(db, nil),
	}
}
