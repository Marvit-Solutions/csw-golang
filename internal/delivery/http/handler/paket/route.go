package paket

import (
	"github.com/gin-gonic/gin"
)

func (paketHandler *PaketHandler) RegisterRoutes(r *gin.RouterGroup) {
	paketGroup := r.Group("/paket")
	paketGroup.GET("/all", paketHandler.ListPaket)
	paketGroup.GET("/top", paketHandler.GetTop3Paket)
	// paketGroup.POST("/", paketHandler.CreatePaket)
	// paketGroup.PUT("/:id_paket", paketHandler.UpdatePaket)
	// paketGroup.DELETE("/:id_paket", paketHandler.DeletePaket)

	// subpaketGroup := r.Group("/subpaket")
	// subpaketGroup.GET("/:id_paket", paketHandler.ListSubPaket)
	// subpaketGroup.POST("/", paketHandler.CreateSubPaket)
	// subpaketGroup.PUT("/:id_subpaket", paketHandler.UpdateSubPaket)
	// subpaketGroup.DELETE("/:id_subpaket", paketHandler.DeleteSubPaket)
	// subpaketGroup.GET("/top", paketHandler.GetTopSubPaket)

}
