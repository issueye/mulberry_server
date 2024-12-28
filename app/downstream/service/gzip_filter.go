package service

import (
	"mulberry/host/app/downstream/model"
	"mulberry/host/app/downstream/requests"
	commonModel "mulberry/host/common/model"
	"mulberry/host/common/service"

	"gorm.io/gorm"
)

type GzipFilter struct {
	service.BaseService[model.GzipFilterInfo]
}

func NewGzipFilter(args ...any) *GzipFilter {
	srv := &GzipFilter{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListGzipFilter
// 根据条件查询列表
func (r *GzipFilter) ListGzipFilter(condition *commonModel.PageQuery[*requests.QueryGzipFilter]) (*commonModel.ResPage[model.GzipFilterInfo], error) {
	return service.GetList[model.GzipFilterInfo](condition, func(qu *requests.QueryGzipFilter, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (r *GzipFilter) UpdateGzipFilterStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
