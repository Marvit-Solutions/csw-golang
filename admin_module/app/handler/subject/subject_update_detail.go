package subjectAdmin

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) SubjectUpdateDetail(ctx *gin.Context) {
	var req request.ParamSubjectAdminUpdateDetail
	if err := helper.ValidateURLParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	if err := helper.ValidateQueryParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	quizAll, err := h.u.SubjectAdminUpdateDetail(req)

	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), quizAll)
}
