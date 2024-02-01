package dto

type SubPlanResponse struct {
	ID                 string `json:"IdSubPlan" form:"IdSubPlan"`
	PlanName           string `json:"PlanName" form:"PlanName"`
	SubPlanName        string `json:"SubPlanName" form:"SubPlanName"`
	SubPlanDescription string `json:"SubPlanDescription" form:"SubPlanDescription"`
	Price              int    `json:"Price" form:"Price"`
}

type SubPlanRequest struct {
	IDPaket            string `json:"IDPaket" form:"IDPaket"`
	SubPlanName        string `json:"SubPlanName" form:"SubPlanName"`
	SubPlanDescription string `json:"SubPlanDescription" form:"SubPlanDescription"`
	Price              int    `json:"Price" form:"Price"`
}

type TopSubPlanResponse struct {
	SubPlanID   string `json:"IdSubPlan" form:"IdSubPlan"`
	SubPlanName string `json:"SubPlanName" form:"SubPlanName"`
	Total       int    `json:"Total" form:"Total"`
}
