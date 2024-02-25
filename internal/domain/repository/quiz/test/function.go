package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
)

func (tr testRepo) GetAllTests(req request.QuizParamRequest) (*[]dto.QuizResponse, error) {
	var quizzes []dto.QuizResponse

	err := tr.db.Table("modules m").
		Select("sttq.id, sttq.test_type, sttq.title, sttq.meeting_date, sttq.open, sttq.description, sttq.time, sttq.point, sttq.attempt").
		Joins("inner join sub_modules sm on sm.module_id = m.id").
		Joins("inner join subjects s on s.sub_module_id = sm.id").
		Joins("inner join subject_test_type_quizzes sttq on sttq.subject_id = s.id").
		Where("m.name LIKE ? and sm.name LIKE ? and sttq.test_type LIKE ?", "%"+req.Module+"%", "%"+req.SubModule+"%", "%"+req.TestType+"%").
		Limit(10).
		Offset(0).
		Find(&quizzes).
		Error
	if err != nil {
		return &quizzes, err
	}

	return &quizzes, nil
}
