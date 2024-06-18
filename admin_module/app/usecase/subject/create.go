package subjectAdmin

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Create(req request.SubjectAdminPayload) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
		"name": req.SubModuleName,
	})

	fmt.Println(req.Name, subModule.ID)
	if err != nil {
		return fmt.Errorf("failed to find subModule: %v", err)
	}
	SubjectData := &model.Subject{
		SubModuleID: subModule.ID,
		Name:        req.Name,
	}
	_, err = u.subjectRepo.Create(SubjectData, tx)
	if err != nil {
		fmt.Println("masuk sini")
		return fmt.Errorf("failed to create quiz: %v", err)
	}
	tx.Commit()

	return nil
}
