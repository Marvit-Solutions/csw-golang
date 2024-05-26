package quiz

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) QuizAll(ctx *gin.Context) {
	var req request.ParamQuizAll
	if err := helper.ValidateURLParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	if err := helper.ValidateQueryParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	quizAll, totalRows, totalPages, err := h.u.QuizAll(req)

	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	meta := helper.Pagination{
		Page:       req.Page,
		Limit:      req.Limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
	}

	helper.NewSuccessResponsePaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), quizAll, meta)
}
