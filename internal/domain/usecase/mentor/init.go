package mentor

import (
	"csw-golang/internal/domain/entity/dto"
	mr "csw-golang/internal/domain/repository/mentor"
)

type MentorUsecase interface {
	GetListTopThreeMentors() (error, dto.ListMentor)
}

type mentorUsecase struct {
	mentorRepo mr.MentorRepo
}

func New(mentorRepo mr.MentorRepo) *mentorUsecase {
	return &mentorUsecase{
		mentorRepo,
	}
}
