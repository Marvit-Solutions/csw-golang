package home

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) PlanAll(req request.PlanHome) ([]*response.PlanHome, error) {
	moduleFilter := req.Module
	if moduleFilter == "" {
		moduleFilter = "skd,matematika"
	}

	modules, err := u.moduleRepo.FindBy(map[string]interface{}{
		"slug": strings.Split(strings.ToLower(moduleFilter), ","),
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	moduleIDs := make([]int, len(modules))
	for i, mod := range modules {
		moduleIDs[i] = mod.ID
	}

	plans, err := u.planRepo.FindBy(map[string]interface{}{
		"module_id": moduleIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find plan: %v", err)
	}

	if req.Name != "" {
		filteredPlans := make([]*model.Plan, 0)
		for _, plan := range plans {
			if strings.Contains(strings.ToLower(plan.Name), strings.ToLower(req.Name)) {
				filteredPlans = append(filteredPlans, plan)
			}
		}
		plans = filteredPlans
	}

	moduleMap := make(map[int]*model.Module)
	for _, mod := range modules {
		moduleMap[mod.ID] = mod
	}

	results := make([]*response.PlanHome, len(plans))
	for i, plan := range plans {
		moduleName := ""
		if mod, ok := moduleMap[plan.ModuleID]; ok {
			moduleName = mod.Name
		}

		results[i] = &response.PlanHome{
			UUID:       plan.UUID,
			ModuleName: moduleName,
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
