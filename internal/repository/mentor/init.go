package mentor

import (
	dto "csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type MentorRepository interface {
	GetListTopThreeMentors() (dto.ListMentor, error)
	GetAllMentors() (dto.ListMentor, error)
}

type mentorRepository struct {
	db *gorm.DB
}

func NewMentorRepository(db *gorm.DB) MentorRepository {
	return &mentorRepository{db}
}
