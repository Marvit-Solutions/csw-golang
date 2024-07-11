package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
)

type MaterialAdmin interface {
	GetMaterialAdminAll(searchKeywords string, subjectID int, subModulID int, limit int, offset int) ([]*response.MaterialAdminAllResponse, error)
	CountMaterialAdminALL(searchKeywords string, subjectID int, subModulID int) (int, error)
}
