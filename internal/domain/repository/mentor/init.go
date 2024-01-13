package mentor

import (
	"csw-golang/internal/domain/entity/dto"
	"gorm.io/gorm"
)

type MentorRepo interface {
	GetListTopThreeMentors() (error, dto.ListMentor)
}

type mentorRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) MentorRepo {
	return &mentorRepo{
		db,
	}
}
