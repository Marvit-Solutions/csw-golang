package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/response"
	"fmt"
)

func (tr testRepo) GetAllTests(req request.QuizRequest) (*[]dto.QuizResponse, *response.Meta, error) {
	var quizzes []dto.QuizResponse
	var count int64
	var perPage int64
	var meta *response.Meta

	err := tr.db.Table("modules m").
		Select("CEIL(COUNT(*)::NUMERIC / 10)").
		Joins("inner join sub_modules sm on sm.module_id = m.id").
		Joins("inner join subjects s on s.sub_module_id = sm.id").
		Joins("inner join subject_test_type_quizzes sttq on sttq.subject_id = s.id").
		Where("m.name LIKE ? AND sm.name LIKE ? AND sttq.test_type LIKE ?", "%"+req.Module+"%", "%"+req.SubModule+"%", "%"+req.TestType+"%").
		Scan(&count).
		Error
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get count: %v", err)
	}

	perPage = req.Limit
	lastPage := count / perPage
	if count%perPage != 0 {
		lastPage++
	}

	err = tr.db.Table("modules m").
		Select("sttq.id, sttq.test_type, sttq.title, sttq.meeting_date, sttq.open, sttq.description, sttq.time, sttq.point, sttq.attempt").
		Joins("inner join sub_modules sm on sm.module_id = m.id").
		Joins("inner join subjects s on s.sub_module_id = sm.id").
		Joins("inner join subject_test_type_quizzes sttq on sttq.subject_id = s.id").
		Where("m.name LIKE ? and sm.name LIKE ? and sttq.test_type LIKE ?", "%"+req.Module+"%", "%"+req.SubModule+"%", "%"+req.TestType+"%").
		Limit(int(req.Limit)).
		Offset(req.Offset).
		Find(&quizzes).
		Error
	if err != nil {
		return &quizzes, nil, fmt.Errorf("failed to get data: %v", err)
	}

	meta = &response.Meta{
		PerPage:   perPage,
		LastPage:  lastPage,
		TotalPage: count,
	}

	return &quizzes, meta, nil
}
