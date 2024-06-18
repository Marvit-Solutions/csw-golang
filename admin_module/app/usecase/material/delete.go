package materialAdmin

import "fmt"

func (u *usecase) Delete(uuid string) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	subSubject, err := u.subSubjectRepo.FindOneBy(map[string]interface{}{
		"uuid": uuid,
	})
	if err != nil {
		return fmt.Errorf("failed to find subSubject: %v", err)
	}

	err = u.subSubjectRepo.Delete(subSubject, tx)

	if err != nil {
		return fmt.Errorf("failed to delete subSubject: %v", err)
	}

	tx.Commit()
	return nil
}
