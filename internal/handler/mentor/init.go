package mentor

import (
	"csw-golang/internal/usecase/mentor"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type mentorHandler struct {
	mentorUsecase mentor.MentorUsecase
}

type MentorHandler interface {
	ListThreeMentors(c *gin.Context)
	GetAllMentors(c *gin.Context)
}

func NewMentorHandler(
	db *gorm.DB,
) MentorHandler {
	return &mentorHandler{
		mentorUsecase: mentor.NewMentorUsecase(db),
	}
}
