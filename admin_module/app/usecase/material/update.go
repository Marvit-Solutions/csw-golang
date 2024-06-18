package materialAdmin

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Update(req request.MaterialAdminPayloadUpdate) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"name": req.SubjectName,
	})

	fmt.Println(req.SubjectName, req.Name, req.Content, subject.ID)
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}

	subSubject, err := u.subSubjectRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})

	if err != nil {
		return fmt.Errorf("failed to find subSubject: %v", err)
	}

	subSubjectData := &model.SubSubject{
		ID:        subSubject.ID,
		UUID:      subSubject.UUID,
		SubjectID: subject.ID,
		Name:      req.Name,
		Content:   req.Content,
	}
	err = u.subSubjectRepo.Update(subSubjectData, tx)
	if err != nil {
		fmt.Println("masuk sini")
		return fmt.Errorf("failed to create quiz: %v", err)
	}
	tx.Commit()

	return nil
}
