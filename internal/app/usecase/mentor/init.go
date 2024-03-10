package mentor

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/repository/mentor"
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type MentorUsecase interface {
	GetListTopThreeMentors() (dto.ListMentor, error)
	GetAllMentors() (dto.ListMentor, error)
}

type mentorUsecase struct {
	mentorRepo mentor.MentorRepository
}

func NewMentorUsecase(
	db *gorm.DB,
) MentorUsecase {
	return &mentorUsecase{
		mentorRepo: mentor.NewMentorRepository(db),
	}
}
