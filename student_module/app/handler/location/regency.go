package location

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/gin-gonic/gin"
)

func (h *handler) Regency(ctx *gin.Context) {
	var req request.LocationRequest
	apiUrl := h.conf.GetString("api-url.binderbyte-regency")

	if err := helper.ValidateURLParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	url := apiUrl + req.ProvinceID
	req.Url = url

	provinces, err := h.u.Regency(req)
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
