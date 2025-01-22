package service

import (
	"mulberry/internal/app/downstream/model"
	"mulberry/internal/app/downstream/requests"
	commonModel "mulberry/internal/common/model"
	"mulberry/internal/common/service"

	"gorm.io/gorm"
)

type Page struct {
	service.BaseService[model.PageInfo]
}

func NewPage(args ...any) *Page {
	srv := &Page{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListPage
// 根据条件查询列表
func (r *Page) ListPage(condition *commonModel.PageQuery[*requests.QueryPage]) (*commonModel.ResPage[model.PageInfo], error) {
	return service.GetList[model.PageInfo](condition, func(qu *requests.QueryPage, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		if qu.Port > 0 {
			d = d.Where("port = ?", qu.Port)
		}

		// 预加载页面版本信息
		d = d.Preload("PageVersion")

		return d
	})
}

func (r *Page) UpdatePageStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
