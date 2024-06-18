package materialAdmin

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Create(req request.MaterialAdminPayload) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"name": req.SubjectName,
	})

	fmt.Println(req.SubjectName, req.Name, req.Content, subject.ID)
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}
	subSubjectData := &model.SubSubject{
		SubjectID: subject.ID,
		Name:      req.Name,
		Content:   req.Content,
	}
	_, err = u.subSubjectRepo.Create(subSubjectData, tx)
	if err != nil {
		fmt.Println("masuk sini")
		return fmt.Errorf("failed to create quiz: %v", err)
	}
	tx.Commit()

	return nil
}
