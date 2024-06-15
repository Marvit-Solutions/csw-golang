package location

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/gin-gonic/gin"
)

func (h *handler) District(ctx *gin.Context) {
	var req request.LocationRequest
	apiUrl := h.conf.GetString("api-url.binderbyte-district")

	if err := helper.ValidateURLParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	url := apiUrl + req.RegencyID
	req.Url = url

	provinces, err := h.u.District(req)
	if err != nil && err == helper.ErrDataNotFound {
		helper.NewErrorResponse(ctx, http.StatusNotFound, http.StatusText(http.StatusNotFound), err.Error())
		return
	}

	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), provinces)
}
