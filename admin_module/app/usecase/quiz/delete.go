package quizAdmin

import "fmt"

func (u *usecase) Delete(uuid string) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": uuid,
	})
	if err != nil {
		return fmt.Errorf("failed to find quiz: %v", err)
	}

	err = u.quizRepo.Delete(quiz, tx)

	if err != nil {
		return fmt.Errorf("failed to delete quiz: %v", err)
	}

	tx.Commit()
	return nil
}
