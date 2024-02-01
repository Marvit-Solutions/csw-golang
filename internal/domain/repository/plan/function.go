package plan

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"sort"
)

func (pr *planRepo) ListPlan() ([]dto.PlanResponse, error) {
	var PlanList []datastruct.Plans

	err := pr.db.Preload("SubPlan").Preload("SubPlan.SubPlanDetail").Find(&PlanList).Error
	if err != nil {
		return nil, err
	}

	var PlanResponses []dto.PlanResponse
	for _, Plan := range PlanList {
		PlanResponse := dto.PlanResponse{
			ID:       Plan.ID,
			PlanName: Plan.Name,
		}

		for _, SubPlan := range Plan.SubPlan {
			SubPlanResponse := dto.SubPlan{
				ID:          SubPlan.ID,
				SubPlanName: SubPlan.Name,
				Price:       SubPlan.Price,
				SubPlanDetail: dto.SubPlanDetail{
					ID:          SubPlan.SubPlanDetail.ID,
					GrupPejuang: SubPlan.SubPlanDetail.GrupPejuang,
					Exercise:    SubPlan.SubPlanDetail.Exercise,
					Access:      SubPlan.SubPlanDetail.Access,
					Module:      SubPlan.SubPlanDetail.Module,
					TryOut:      SubPlan.SubPlanDetail.TryOut,
					Zoom:        SubPlan.SubPlanDetail.Zoom,
				},
			}

			PlanResponse.SubPlan = append(PlanResponse.SubPlan, SubPlanResponse)
		}

		PlanResponses = append(PlanResponses, PlanResponse)
	}

	return PlanResponses, nil
}

func (pr *planRepo) GetTop3Plan() ([]dto.SubPlanTop3Response, error) {
	var PlanList []datastruct.Plans

	err := pr.db.Preload("SubPlan").Preload("SubPlan.SubPlanDetail").Find(&PlanList).Error
	if err != nil {
		return nil, err
	}

	var PlanResponses []dto.SubPlanTop3Response
	for _, Plan := range PlanList {
		for _, SubPlan := range Plan.SubPlan {
			response := dto.SubPlanTop3Response{
				ID:          Plan.ID,
				PlanName:    Plan.Name,
				SubPlanName: SubPlan.Name,
				Price:       SubPlan.Price,
				GrupPejuang: SubPlan.SubPlanDetail.GrupPejuang,
				Exercise:    SubPlan.SubPlanDetail.Exercise,
				Access:      SubPlan.SubPlanDetail.Access,
				Module:      SubPlan.SubPlanDetail.Module,
				TryOut:      SubPlan.SubPlanDetail.TryOut,
				Zoom:        SubPlan.SubPlanDetail.Zoom,
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
