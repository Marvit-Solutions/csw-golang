package quiz

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/gin-gonic/gin"
)

func (h *handler) QuizReview(ctx *gin.Context) {
	var req request.ParamQuizReview

	if err := helper.ValidateURLParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	authenticatedUser, err := auth.GetAuthenticatedUser(ctx.Request)
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), err.Error())
		return
	}

	req.AuthenticatedUser = authenticatedUser

	quizzes, err := h.u.QuizReview(req)
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), quizzes)
}
