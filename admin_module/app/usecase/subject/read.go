package subjectAdmin

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Read() ([]*model.Subject, error) {
	subjects, err := u.subjectRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find subjects: %v", err)
	}
	result := subjects
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID > result[j].ID
	})
	return subjects, nil
}
