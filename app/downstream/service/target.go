package service

import (
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	commonModel "mulberry/common/model"
	"mulberry/common/service"

	"gorm.io/gorm"
)

type Target struct {
	service.BaseService[model.TargetInfo]
}

func NewTarget(args ...any) *Target {
	srv := &Target{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListTarget
// 根据条件查询列表
func (r *Target) ListTarget(condition *commonModel.PageQuery[*requests.QueryTarget]) (*commonModel.ResPage[model.TargetInfo], error) {
	return service.GetList[model.TargetInfo](condition, func(qu *requests.QueryTarget, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (r *Target) UpdateTargetStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
