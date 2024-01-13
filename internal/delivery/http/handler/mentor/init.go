package mentor

import (
	mc "csw-golang/internal/domain/usecase/mentor"
)

type MentorHandler struct {
	mentorUsecase mc.MentorUsecase
}

func New(mentorUsecase mc.MentorUsecase) *MentorHandler {
	return &MentorHandler{
		mentorUsecase,
	}
}
