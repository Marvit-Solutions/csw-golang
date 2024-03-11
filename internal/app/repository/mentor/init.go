package mentor

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type MentorRepository interface {
	GetListTopThreeMentors() (response.ListMentor, error)
	GetAllMentors() (response.ListMentor, error)
}

type mentorRepository struct {
	db *gorm.DB
}

func NewMentorRepository(db *gorm.DB) MentorRepository {
	return &mentorRepository{db}
}
