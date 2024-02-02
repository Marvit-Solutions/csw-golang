package modul

// func (mh *ModulHandler) ListModul(c *gin.Context) {

// 	response, err := mh.modulUsecase.ListModul()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Fail{
// 			Code:    http.StatusInternalServerError,
// 			Status:  http.StatusText(http.StatusInternalServerError),
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if len(response) == 0 {
// 		c.JSON(http.StatusNotFound, dto.Success[interface{}]{
// 			Message: "Modul tidak ditemukan!",
// 			Code:    http.StatusNotFound,
// 			Status:  http.StatusText(http.StatusNotFound),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.Success[interface{}]{
// 		Message: "Berhasil mendapatkan Modul!",
// 		Code:    http.StatusOK,
// 		Status:  http.StatusText(http.StatusOK),
// 		Data:    response,
// 	})
// }
