package service

import (
	"mulberry/internal/app/downstream/model"
	"mulberry/internal/app/downstream/requests"
	commonModel "mulberry/internal/common/model"
	"mulberry/internal/common/service"

	"gorm.io/gorm"
)

type PageVersion struct {
	service.BaseService[model.PageVersionInfo]
}

func NewPageVersion(args ...any) *PageVersion {
	srv := &PageVersion{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListPage
// 根据条件查询列表
func (r *PageVersion) ListPageVersion(condition *commonModel.PageQuery[*requests.QueryPageVersion]) (*commonModel.ResPage[model.PageVersionInfo], error) {
	return service.GetList[model.PageVersionInfo](condition, func(qu *requests.QueryPageVersion, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}
