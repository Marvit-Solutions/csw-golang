package dto

type SubPlanDetail struct {
	ID          string `json:"ID" form:"ID"`
	GrupPejuang bool   `json:"GrupPejuang" form:"GrupPejuang"`
	Exercise    uint   `json:"Exercise" form:"Exercise"`
	Access      uint   `json:"Access" form:"Access"`
	Module      bool   `json:"Module" form:"Module"`
	TryOut      uint   `json:"TryOut" form:"TryOut"`
	Zoom        bool   `json:"Zoom" form:"Zoom"`
}

type SubPlan struct {
	ID            string        `json:"ID" form:"ID"`
	SubPlanName   string        `json:"SubPlanName" form:"SubPlanName"`
	Price         float64       `json:"Price" form:"Price"`
	SubPlanDetail SubPlanDetail `json:"SubPlanDetail" form:"SubPlanDetail"`
}
type PlanResponse struct {
	ID       string    `json:"ID" form:"ID"`
	PlanName string    `json:"PlanName" form:"PlanName"`
	SubPlan  []SubPlan `json:"SubPlan" form:"SubPlan"`
}

type SubPlanTop3Response struct {
	ID          string  `json:"ID" form:"ID"`
	PlanName    string  `json:"PlanName" form:"PlanName"`
	SubPlanName string  `json:"SubPlanName" form:"SubPlanName"`
	Price       float64 `json:"Price" form:"Price"`
	GrupPejuang bool    `json:"GrupPejuang" form:"GrupPejuang"`
	Exercise    uint    `json:"Exercise" form:"Exercise"`
	Access      uint    `json:"Access" form:"Access"`
	Module      bool    `json:"Module" form:"Module"`
	TryOut      uint    `json:"TryOut" form:"TryOut"`
	Zoom        bool    `json:"Zoom" form:"Zoom"`
}

type PaketRequest struct {
	PlanName        string `json:"PlanName" form:"PlanName"`
	PlanDescription string `json:"PlanDescription" form:"PlanDescription"`
}
