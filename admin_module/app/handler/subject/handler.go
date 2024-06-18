package subjectAdmin

import (
	subjectAdmin "github.com/Marvit-Solutions/csw-golang/admin_module/app/usecase/subject"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u subjectAdmin.Usecase
}

type Handler interface {
	SubjectAdminAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	SubjectUpdateDetail(c *gin.Context)
	Read(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: subjectAdmin.NewUsecase(db),
	}
}
