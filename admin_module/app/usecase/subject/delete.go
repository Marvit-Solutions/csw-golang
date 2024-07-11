package subjectAdmin

import "fmt"

func (u *usecase) Delete(uuid string) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"uuid": uuid,
	})
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}

	err = u.subjectRepo.Delete(subject, tx)

	if err != nil {
		return fmt.Errorf("failed to delete subject: %v", err)
	}

	tx.Commit()
	return nil
}
