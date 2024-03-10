package mentor

import (
	dto "csw-golang/internal/domain/response"
	"csw-golang/internal/repository/mentor"

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
