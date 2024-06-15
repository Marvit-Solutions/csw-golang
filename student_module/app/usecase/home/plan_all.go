package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

func (u *usecase) PlanAll(req request.PlanHome) ([]*response.PlanHome, error) {

	plans, planMediaIDs, err := u.localHomeRepo.FindPlanInfo(req, false)
	if err != nil {
		return nil, err
	}

	mediaMaps, err := u.localHomeRepo.FindMediaInfo(planMediaIDs)
	if err != nil {
		return nil, err
	}

	mediaMapPlan, mediaMentorFound := mediaMaps["media1"]
	if !mediaMentorFound {
		return nil, fmt.Errorf("failed to find  media mentor")
	}

	results := make([]*response.PlanHome, 0)
	for _, plan := range plans {
		results = append(results, &response.PlanHome{
			UUID:       plan.UUID,
			ModuleName: plan.ModuleName,
			Name:       plan.Name,
			Media:      helper.MultiResImages(mediaMapPlan[plan.MediaID]),
			Price:      plan.Price,
			Group:      plan.Group,
			Exercise:   plan.Exercise,
			Access:     plan.Access,
			Module:     plan.Module,
			TryOut:     plan.TryOut,
			Zoom:       plan.Zoom,
		})
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	return results, nil
}
