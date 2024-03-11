package plan

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/datastruct"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (pr *planRepository) ListPlan() ([]response.PlanResponse, error) {
	var PlanList []datastruct.Modules

	err := pr.db.Preload("SubPlans").Find(&PlanList).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get plans: %v", err)
	}

	var data []response.PlanResponse
	for _, Plan := range PlanList {
		PlanResponse := response.PlanResponse{
			ID:       Plan.ID,
			PlanName: Plan.Name,
		}

		for _, SubPlan := range Plan.SubPlans {
			SubPlanResponse := response.SubPlan{
				ID:          SubPlan.ID,
				SubPlanName: SubPlan.Name,
				Price:       SubPlan.Price,
				GrupPejuang: SubPlan.GrupPejuang,
				Exercise:    SubPlan.Exercise,
				Access:      SubPlan.Access,
				Module:      SubPlan.Module,
				TryOut:      SubPlan.TryOut,
				Zoom:        SubPlan.Zoom,
			}

			PlanResponse.SubPlan = append(PlanResponse.SubPlan, SubPlanResponse)
		}

		data = append(data, PlanResponse)
	}

	return data, nil
}

func (pr *planRepository) GetTop3Plan() ([]response.SubPlanTop3Response, error) {
	var PlanList []datastruct.Modules

	err := pr.db.Preload("SubPlans").Find(&PlanList).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get plans: %v", err)
	}

	var data []response.SubPlanTop3Response
	for _, Plan := range PlanList {
		for _, SubPlan := range Plan.SubPlans {
			response := response.SubPlanTop3Response{
				ID:          Plan.ID,
				PlanName:    Plan.Name,
				SubPlanName: SubPlan.Name,
				Price:       SubPlan.Price,
				GrupPejuang: SubPlan.GrupPejuang,
				Exercise:    SubPlan.Exercise,
				Access:      SubPlan.Access,
				Module:      SubPlan.Module,
				TryOut:      SubPlan.TryOut,
				Zoom:        SubPlan.Zoom,
			}

			data = append(data, response)
		}
	}

	// Sorting descending berdasarkan Price (Belum berdasarkan rating pembeli)
	// digabung, antara skd dan mtk
	sort.Slice(data, func(i, j int) bool {
		return data[i].Price > data[j].Price
	})

	if len(data) > 3 {
		data = data[:3]
	}

	return data, nil
}
