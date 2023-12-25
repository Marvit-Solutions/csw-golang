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
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Code:    http.StatusOK,
		Message: "List Paket Success",
		Data:    response,
	})
}

func (ph *PaketHandler) ListSubPaket(c *gin.Context) {
	idPaket := c.Param("id_paket")
	if idPaket == "" {
		c.JSON(http.StatusBadRequest, dto.Fail{
			Code:    http.StatusBadRequest,
			Message: "ID Paket is empty",
		})
		return
	}

	response, err := ph.paketUsecase.ListSubPaket(idPaket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

		c.JSON(http.StatusOK, dto.Success[interface{}]{
			Code:    http.StatusOK,
			Message: "List Subpaket Success",
			Data:    response,
		})
	}
}
