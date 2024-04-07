package home

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) PlanAll() ([]*response.PlanHome, error) {
	plans, err := u.planRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find plan: %v", err)
	}

	planIDs := make([]int, len(plans))
	for i, plan := range plans {
		planIDs[i] = plan.ModuleID
	}

	modules, err := u.moduleRepo.FindBy(map[string]interface{}{
		"id": planIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	moduleMap := make(map[int]*model.Module)
	for _, module := range modules {
		moduleMap[module.ID] = module
	}

	results := make([]*response.PlanHome, len(plans))
	for i, plan := range plans {
		results[i] = &response.PlanHome{
			UUID:       plan.UUID,
			ModuleName: moduleMap[plan.ModuleID].Name,
			Name:       plan.Name,
			Picture:    plan.Picture,
			Price:      plan.Price,
			Group:      plan.Group,
			Exercise:   int(plan.Exercise),
			Access:     int(plan.Access),
			Module:     plan.Module,
			TryOut:     int(plan.TryOut),
			Zoom:       plan.Zoom,
		}
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Price < results[j].Price
	})

	return results, nil
}
