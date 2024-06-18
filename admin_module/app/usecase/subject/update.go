package subjectAdmin

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Update(req request.SubjectAdminPayloadUpdate) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})

	fmt.Println(req.Name, subject.ID)
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}

	subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
		"id": subject.SubModuleID,
	})

	if err != nil {
		return fmt.Errorf("failed to find subModule: %v", err)
	}

	subjectData := &model.Subject{
		ID:          subject.ID,
		UUID:        subject.UUID,
		SubModuleID: subModule.ID,
		Name:        req.Name,
	}
	err = u.subjectRepo.Update(subjectData, tx)
	if err != nil {
		fmt.Println("masuk sini")
		return fmt.Errorf("failed to create subjectData: %v", err)
	}
	tx.Commit()

	return nil
}
