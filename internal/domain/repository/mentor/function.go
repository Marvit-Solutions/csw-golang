package mentor

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (m mentorRepo) GetListTopThreeMentors() (error, dto.ListMentor) {
	var topMentors dto.ListMentor
	var mentors []datastruct.Mentor
	_ = m.db.Order("rating desc, name").Limit(3).Find(&mentors)

	// Convert the database mentors to DTO mentors
	for _, mentor := range mentors {
		dtoMentor := struct {
			Id           string  `json:"id"`
			Name         string  `json:"name"`
			Description  string  `json:"description"`
			ProfilePhoto string  `json:"profile_photo"`
			Rating       float32 `json:"rating"`
		}{
			Id:           mentor.Id,
			Name:         mentor.Name,
			Description:  mentor.Description,
			ProfilePhoto: mentor.ProfilePhoto,
			Rating:       mentor.Rating,
			// Add other fields as needed
		}
		topMentors = append(topMentors, dtoMentor)
	}
	return nil, topMentors
}
