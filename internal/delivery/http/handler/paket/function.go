package paket

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ph *PaketHandler) ListPaket(c *gin.Context) {

	response, err := ph.paketUsecase.ListPaket()
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
			Message: "Paket tidak ditemukan!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Berhasil mendapatkan paket!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}

// func (ph *PaketHandler) CreatePaket(c *gin.Context) {
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

// func (ph *PaketHandler) UpdatePaket(c *gin.Context) {
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

// // func (ph *PaketHandler) DeletePaket(c *gin.Context) {
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

// // 	if response == (dto.PaketResponse{}) {
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

// func (ph *PaketHandler) ListSubPaket(c *gin.Context) {
// 	idPaket := c.Param("id_paket")
// 	if idPaket == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID Paket is empty",
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.ListSubPaket(idPaket)
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
// 			Message: "Subpaket not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "List Subpaket Success",
// 		Data:    response,
// 	})
// }

// func (ph *PaketHandler) CreateSubPaket(c *gin.Context) {
// 	var request dto.SubPaketRequest

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.CreateSubPaket(request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Create Subpaket Success",
// 		Data:    response,
// 	})
// }

// func (ph *PaketHandler) UpdateSubPaket(c *gin.Context) {
// 	var request dto.SubPaketRequest
// 	idSubPaket := c.Param("id_sub_paket")
// 	if idSubPaket == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID Subpaket is empty",
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

// 	response, err := ph.paketUsecase.UpdateSubPaket(request, idSubPaket)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Update Subpaket Success",
// 		Data:    response,
// 	})
// }

// func (ph *PaketHandler) DeleteSubPaket(c *gin.Context) {
// 	idSubPaket := c.Param("id_sub_paket")
// 	if idSubPaket == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID Subpaket is empty",
// 		})
// 		return
// 	}

// 	response, err := ph.paketUsecase.DeleteSubPaket(idSubPaket)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if response == (dto.SubPaketResponse{}) {
// 		c.JSON(http.StatusNotFound, dto.Fail{
// 			Code:    http.StatusNotFound,
// 			Message: "Subpaket not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Delete Subpaket Success",
// 		Data:    response,
// 	})
// }

// func (ph *PaketHandler) GetTopSubPaket(c *gin.Context) {
// 	response, err := ph.paketUsecase.GetTopSubPaket()
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
// 			Message: "Subpaket not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Get Top Subpaket Success",
// 		Data:    response,
// 	})
// }
