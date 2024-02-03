package module

import (
	"csw-golang/internal/domain/entity/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *ModuleHandler) GetListModules(c *gin.Context) {
	fmt.Println("GetListModules")
	response, err := mc.moduleUsecase.GetListModules()
	fmt.Println("response", response)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})

}
