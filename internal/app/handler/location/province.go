package location

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) Province(ctx *gin.Context) {
	var req request.LocationRequest
	apiKey := h.conf.GetString("api-key.binderbyte")
	apiUrl := h.conf.GetString("api-url.binderbyte")
	url := apiUrl + "provinsi?api_key=" + apiKey
	req.Url = url

	provinces, err := h.u.Province(req)
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
