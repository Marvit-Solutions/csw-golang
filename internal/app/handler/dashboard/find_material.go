package dashboard

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
	"github.com/gin-gonic/gin"
)

func (h *handler) FindMaterial(c *gin.Context) {
	var req request.ParamDashboard

	authenticatedUser, err := auth.GetAuthenticatedUser(c.Request)
	if err != nil {
		helper.NewErrorResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), err.Error())
		return
	}

	req.AuthenticatedUser = authenticatedUser

	materials, err := h.u.FindMaterial(req)
	if err != nil && err == helper.ErrAccessDenied {
		helper.NewErrorResponse(c, http.StatusForbidden, http.StatusText(http.StatusForbidden), err.Error())
		return
	}
	if err != nil {
		helper.NewErrorResponse(c, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), materials)
}
