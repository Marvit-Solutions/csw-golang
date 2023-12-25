package paket

import (
	"github.com/gin-gonic/gin"
)

func (paketHandler *PaketHandler) RegisterRoutes(r *gin.RouterGroup) {
	paketGroup := r.Group("/paket")
	paketGroup.GET("/", paketHandler.ListPaket)

	subpaketGroup := r.Group("/subpaket")
	subpaketGroup.GET("/:id_paket", paketHandler.ListSubPaket)

}
