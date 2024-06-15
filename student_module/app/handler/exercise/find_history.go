package exercise

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/gin-gonic/gin"
)

func (h *handler) FindHistory(c *gin.Context) {
	var req request.ExerciseHistory

	if err := helper.ValidateURLParams(c, &req); err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	authenticatedUser, err := auth.GetAuthenticatedUser(c.Request)
	if err != nil {
		helper.NewErrorResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), err.Error())
		return
	}

	req.AuthenticatedUser = authenticatedUser

	history, err := h.u.FindHistory(req)
	if err != nil && err == helper.ErrAccessDenied {
		helper.NewErrorResponse(c, http.StatusForbidden, http.StatusText(http.StatusForbidden), err.Error())
		return
	}
	if err != nil {
		helper.NewErrorResponse(c, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), history)
}
