package plan

import (
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/datastruct"
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (pr *planRepository) ListPlan() ([]dto.PlanResponse, error) {
	var PlanList []datastruct.Modules

	err := pr.db.Preload("SubPlans").Find(&PlanList).Error
	if err != nil {
		return nil, err
	}

	var PlanResponses []dto.PlanResponse
	for _, Plan := range PlanList {
		PlanResponse := dto.PlanResponse{
			ID:       Plan.ID,
			PlanName: Plan.Name,
		}

		for _, SubPlan := range Plan.SubPlans {
			SubPlanResponse := dto.SubPlan{
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

		PlanResponses = append(PlanResponses, PlanResponse)
	}

	return PlanResponses, nil
}

func (pr *planRepository) GetTop3Plan() ([]dto.SubPlanTop3Response, error) {
	var PlanList []datastruct.Modules

	err := pr.db.Preload("SubPlans").Find(&PlanList).Error
	if err != nil {
		return nil, err
	}

	var PlanResponses []dto.SubPlanTop3Response
	for _, Plan := range PlanList {
		for _, SubPlan := range Plan.SubPlans {
			response := dto.SubPlanTop3Response{
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

			PlanResponses = append(PlanResponses, response)
		}
	}

	// Sorting descending berdasarkan Price (Belum berdasarkan rating pembeli)
	// digabung, antara skd dan mtk
	sort.Slice(PlanResponses, func(i, j int) bool {
		return PlanResponses[i].Price > PlanResponses[j].Price
	})

	if len(PlanResponses) > 3 {
		PlanResponses = PlanResponses[:3]
	}

	return PlanResponses, nil
}
