package quiz

import (
	"fmt"
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) QuizDetail(ctx *gin.Context) {
	var req request.ParamQuizDetail
	fmt.Println("ini satu")
	if err := helper.ValidateURLParams(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	fmt.Println("ini dua")

	quizDetail, err := h.u.QuizDetail(req)

	fmt.Println("ini quiz detail")
	fmt.Println(quizDetail)
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	fmt.Println("ini tiga")

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), quizDetail)
}
