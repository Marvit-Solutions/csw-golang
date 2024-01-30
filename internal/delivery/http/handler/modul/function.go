package modul

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mh *ModulHandler) ListModul(c *gin.Context) {

	response, err := mh.modulUsecase.ListModul()
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
			Message: "Modul tidak ditemukan!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Berhasil mendapatkan Modul!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}

func (mh *ModulHandler) GetTop3Modul(c *gin.Context) {
	response, err := mh.modulUsecase.GetTop3Modul()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(response) == 0 {
		c.JSON(http.StatusNotFound, dto.Success[interface{}]{
			Message: "Modul tidak ditemukan!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Berhasil mendapatkan Modul!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}

// func (mh *ModulHandler) CreateModul(c *gin.Context) {
// 	var request dto.ModulRequest

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response, err := mh.modulUsecase.CreateModul(request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Create Modul Success",
// 		Data:    response,
// 	})

// }

// func (mh *ModulHandler) UpdateModul(c *gin.Context) {
// 	var request dto.ModulRequest
// 	idModul := c.Param("id_Modul")
// 	if idModul == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID Modul is empty",
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

// 	response, err := mh.modulUsecase.UpdateModul(request, idModul)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Update Modul Success",
// 		Data:    response,
// 	})

// }

// // func (mh *ModulHandler) DeleteModul(c *gin.Context) {
// // 	idModul := c.Param("id_Modul")
// // 	if idModul == "" {
// // 		c.JSON(http.StatusBadRequest, dto.Fail{
// // 			Code:    http.StatusBadRequest,
// // 			Message: "ID Modul is empty",
// // 		})
// // 		return
// // 	}

// // 	response, err := mh.modulUsecase.DeleteModul(idModul)
// // 	if err != nil {
// // 		c.JSON(http.StatusInternalServerError, dto.Fail{
// // 			Code:    http.StatusInternalServerError,
// // 			Message: err.Error(),
// // 		})
// // 		return
// // 	}

// // 	if response == (dto.ModulResponse{}) {
// // 		c.JSON(http.StatusNotFound, dto.Fail{
// // 			Code:    http.StatusNotFound,
// // 			Message: "Modul not found",
// // 		})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// // 		Code:    http.StatusOK,
// // 		Message: "Delete Modul Success",
// // 		Data:    response,
// // 	})
// // }

// func (mh *ModulHandler) ListSubModul(c *gin.Context) {
// 	idModul := c.Param("id_Modul")
// 	if idModul == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID Modul is empty",
// 		})
// 		return
// 	}

// 	response, err := mh.modulUsecase.ListSubModul(idModul)
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
// 			Message: "SubModul not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "List SubModul Success",
// 		Data:    response,
// 	})
// }

// func (mh *ModulHandler) CreateSubModul(c *gin.Context) {
// 	var request dto.SubModulRequest

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response, err := mh.modulUsecase.CreateSubModul(request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Create SubModul Success",
// 		Data:    response,
// 	})
// }

// func (mh *ModulHandler) UpdateSubModul(c *gin.Context) {
// 	var request dto.SubModulRequest
// 	idSubModul := c.Param("id_sub_Modul")
// 	if idSubModul == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID SubModul is empty",
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

// 	response, err := mh.modulUsecase.UpdateSubModul(request, idSubModul)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Update SubModul Success",
// 		Data:    response,
// 	})
// }

// func (mh *ModulHandler) DeleteSubModul(c *gin.Context) {
// 	idSubModul := c.Param("id_sub_Modul")
// 	if idSubModul == "" {
// 		c.JSON(http.StatusBadRequest, dto.Fail{
// 			Code:    http.StatusBadRequest,
// 			Message: "ID SubModul is empty",
// 		})
// 		return
// 	}

// 	response, err := mh.modulUsecase.DeleteSubModul(idSubModul)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if response == (dto.SubModulResponse{}) {
// 		c.JSON(http.StatusNotFound, dto.Fail{
// 			Code:    http.StatusNotFound,
// 			Message: "SubModul not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Delete SubModul Success",
// 		Data:    response,
// 	})
// }

// func (mh *ModulHandler) GetTopSubModul(c *gin.Context) {
// 	response, err := mh.modulUsecase.GetTopSubModul()
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
// 			Message: "SubModul not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Code:    http.StatusOK,
// 		Message: "Get Top SubModul Success",
// 		Data:    response,
// 	})
// }
