package plan

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"sort"
)

func (pr *planRepo) ListPlan() ([]dto.PlanResponse, error) {
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

func (pr *planRepo) GetTop3Plan() ([]dto.SubPlanTop3Response, error) {
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

// func (pr *planRepo) CreatePlan(request dto.PlanRequest) (dto.PlanResponse, error) {
// 	var Plan dto.PlanResponse

// 	err := pr.db.Raw("INSERT INTO Plans (Name_Plan, deskripsi_Plan) VALUES (?, ?) RETURNING id,Name_Plan,deskripsi_Plan", request.PlanName, request.DeskripsiPlan).Scan(&Plan).Error

// 	if err != nil {
// 		return Plan, err
// 	}

// 	return Plan, nil

// }

// func (pr *planRepo) UpdatePlan(request dto.PlanRequest, id string) (dto.PlanResponse, error) {
// 	var Plan dto.PlanResponse

// 	err := pr.db.Raw("UPDATE Plans SET Name_Plan = ?, deskripsi_Plan = ? WHERE id = ? RETURNING id,Name_Plan,deskripsi_Plan", request.PlanName, request.DeskripsiPlan, id).Scan(&Plan).Error

// 	if err != nil {
// 		return Plan, err
// 	}

// 	return Plan, nil

// }

// func (pr *planRepo) DeletePlan(id string) (dto.PlanResponse, error) {
// 	var Plan dto.PlanResponse

// 	err := pr.db.Raw("DELETE FROM Plans WHERE id = ? RETURNING id,Name_Plan,deskripsi_Plan", id).Scan(&Plan).Error
// 	if err != nil {
// 		return Plan, err
// 	}

// 	return Plan, nil
// }

// func (pr *planRepo) ListSubPlan(idPlan string) ([]dto.SubPlanResponse, error) {
// 	var listSubPlan []dto.SubPlanResponse

// 	err := pr.db.Raw("SELECT s.id, p.Name_Plan, s.Name_sub_Plan, s.deskripsi_sub_Plan, s.Price FROM sub_Plans s, Plans p WHERE s.id_Plan = p.id AND id_Plan = ?", idPlan).Scan(&listSubPlan).Error

// 	if err != nil {
// 		return listSubPlan, err
// 	}

// 	return listSubPlan, nil
// }

// func (pr *planRepo) CreateSubPlan(request dto.SubPlanRequest) (dto.SubPlanResponse, error) {
// 	var SubPlan dto.SubPlanResponse

// 	err := pr.db.Raw("INSERT INTO sub_Plans (id_Plan, Name_sub_Plan, deskripsi_sub_Plan, Price) VALUES (?, ?, ?, ?) RETURNING id, Name_sub_Plan, deskripsi_sub_Plan, Price", request.IDPlan, request.SubPlanName, request.DeskripsiSubPlan, request.Price).Scan(&SubPlan).Error

// 	if err != nil {
// 		return SubPlan, err
// 	}

// 	return SubPlan, nil
// }

// func (pr *planRepo) UpdateSubPlan(request dto.SubPlanRequest, id string) (dto.SubPlanResponse, error) {
// 	var SubPlan dto.SubPlanResponse

// 	err := pr.db.Raw("UPDATE sub_Plans SET Name_sub_Plan = ?, deskripsi_sub_Plan = ?, Price = ? WHERE id = ? RETURNING id, Name_sub_Plan, deskripsi_sub_Plan, Price", request.SubPlanName, request.DeskripsiSubPlan, request.Price, id).Scan(&SubPlan).Error

// 	if err != nil {
// 		return SubPlan, err
// 	}

// 	return SubPlan, nil
// }

// func (pr *planRepo) DeleteSubPlan(id string) (dto.SubPlanResponse, error) {
// 	var SubPlan dto.SubPlanResponse

// 	err := pr.db.Raw("DELETE FROM sub_Plans WHERE id = ? RETURNING id, Name_sub_Plan, deskripsi_sub_Plan, Price", id).Scan(&SubPlan).Error

// 	if err != nil {
// 		return SubPlan, err
// 	}

// 	return SubPlan, nil
// }

// func (pr *planRepo) GetTopSubPlan() ([]dto.TopSubPlanResponse, error) {
// 	var ListTopSubPlan []dto.TopSubPlanResponse

// 	err := pr.db.Raw("SELECT s.sub_Plan_id, sp.Name_sub_Plan , COUNT(s.id) AS Total FROM subscriptions s, sub_Plans sp WHERE s.sub_Plan_id = sp.id GROUP BY s.sub_Plan_id, sp.Name_sub_Plan ORDER BY Total DESC LIMIT 3").Scan(&ListTopSubPlan).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return ListTopSubPlan, nil
// }
