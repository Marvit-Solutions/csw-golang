package mentor

import (
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
	"github.com/Marvit-Solutions/csw-golang/internal/repository/mentor"

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
