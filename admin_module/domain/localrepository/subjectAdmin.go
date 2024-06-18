package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
)

type SubjectAdmin interface {
	GetSubjectAdminAll(searchKeywords string, subModulID int, limit int, offset int) ([]*response.SubjectAdminAllResponse, error)
	CountSubjectAdminALL(searchKeywords string, subModulID int) (int, error)
}
