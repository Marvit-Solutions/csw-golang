package paket

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ph *PlanHandler) ListPlan(c *gin.Context) {

	response, err := ph.paketUsecase.ListPlan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	if len(response) == 0 {
		c.JSON(http.StatusNotFound, dto.Success[interface{}]{
			Message: "Plans not found!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success get plans!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}

func (ph *PlanHandler) GetTop3Plan(c *gin.Context) {
	response, err := ph.paketUsecase.GetTop3Plan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(response) == 0 {
		c.JSON(http.StatusNotFound, dto.Success[interface{}]{
			Message: "Plans not found!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success get top 3 plans!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}

// func (ph *PlanHandler) CreatePaket(c *gin.Context) {
// 	var request dto.PaketRequest

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.CreatePaket(request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Create Paket Success",
// 		Data:    response,
// 	})

// }

// func (ph *PlanHandler) UpdatePaket(c *gin.Context) {
// 	var request dto.PaketRequest
// 	idPaket := c.Param("id_paket")
// 	if idPaket == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID Paket is empty",
// 		})
// 		return
// 	}

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.UpdatePaket(request, idPaket)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Update Paket Success",
// 		Data:    response,
// 	})

// }

// // func (ph *PlanHandler) DeletePaket(c *gin.Context) {
// // 	idPaket := c.Param("id_paket")
// // 	if idPaket == "" {
// // 		c.JSON(http.StatusBadRequest, dto.Fail{
// // 			Code:    http.StatusBadRequest,
// // 			Message: "ID Paket is empty",
// // 		})
// // 		return
// // 	}

// // 	response, err := ph.paketUsecase.DeletePaket(idPaket)
// // 	if err != nil {
// // 		c.JSON(http.StatusInternalServerError, dto.Fail{
// // 			Code:    http.StatusInternalServerError,
// // 			Message: err.Error(),
// // 		})
// // 		return
// // 	}

// // 	if response == (dto.PlanResponse{}) {
// // 		c.JSON(http.StatusNotFound, dto.Fail{
// // 			Code:    http.StatusNotFound,
// // 			Message: "Paket not found",
// // 		})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// // 		Code:    http.StatusOK,
// // 		Message: "Delete Paket Success",
// // 		Data:    response,
// // 	})
// // }

// func (ph *PlanHandler) ListSubPlan(c *gin.Context) {
// 	idPaket := c.Param("id_paket")
// 	if idPaket == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID Paket is empty",
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.ListSubPlan(idPaket)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}
// 	if response == nil {
// 		c.JSON(http.StatusNotFound, dto.Fail{
// 			Code:    http.StatusNotFound,
// 			Message: "SubPlan not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "List SubPlan Success",
// 		Data:    response,
// 	})
// }

// func (ph *PlanHandler) CreateSubPlan(c *gin.Context) {
// 	var request dto.SubPlanRequest

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.CreateSubPlan(request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Create SubPlan Success",
// 		Data:    response,
// 	})
// }

// func (ph *PlanHandler) UpdateSubPlan(c *gin.Context) {
// 	var request dto.SubPlanRequest
// 	idSubPlan := c.Param("id_sub_paket")
// 	if idSubPlan == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID SubPlan is empty",
// 		})
// 		return
// 	}

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.UpdateSubPlan(request, idSubPlan)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Update SubPlan Success",
// 		Data:    response,
// 	})
// }

// func (ph *PlanHandler) DeleteSubPlan(c *gin.Context) {
// 	idSubPlan := c.Param("id_sub_paket")
// 	if idSubPlan == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID SubPlan is empty",
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.DeleteSubPlan(idSubPlan)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if response == (dto.SubPlanResponse{}) {
// 		c.JSON(http.StatusNotFound, dto.Fail{
// 			Code:    http.StatusNotFound,
// 			Message: "SubPlan not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Delete SubPlan Success",
// 		Data:    response,
// 	})
// }

// func (ph *PlanHandler) GetTopSubPlan(c *gin.Context) {
// 	response, err := ph.paketUsecase.GetTopSubPlan()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if response == nil {
// 		c.JSON(http.StatusNotFound, dto.Fail{
// 			Code:    http.StatusNotFound,
// 			Message: "SubPlan not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Get Top SubPlan Success",
// 		Data:    response,
// 	})
// }
