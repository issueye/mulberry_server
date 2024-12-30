package service

import (
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	commonModel "mulberry/common/model"
	"mulberry/common/service"

	"gorm.io/gorm"
)

type Rule struct {
	service.BaseService[model.RuleInfo]
}

func NewRule(args ...any) *Rule {
	srv := &Rule{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListRule
// 根据条件查询列表
func (r *Rule) ListRule(condition *commonModel.PageQuery[*requests.QueryRule]) (*commonModel.ResPage[model.RuleInfo], error) {
	return service.GetList[model.RuleInfo](condition, func(qu *requests.QueryRule, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		if qu.Port > 0 {
			d = d.Where("port = ?", qu.Port)
		}

		// 预加载
		d = d.Preload("Target")

		d = d.Order("[order]")

		return d
	})
}

func (r *Rule) GetDatas(condition *requests.QueryRuleDatas) ([]*model.RuleInfo, error) {
	return service.GetDatas[model.RuleInfo](condition, func(qu *requests.QueryRuleDatas, d *gorm.DB) *gorm.DB {
		if qu.Port > 0 {
			d = d.Where("port =?", qu.Port)
		}

		if qu.Status {
			d = d.Where("status = ? ", 1)
		}

		// 预加载
		d = d.Preload("Target")

		// 排序
		d = d.Order("[order]")

		return d
	})
}

func (r *Rule) UpdateRuleStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
