package logic

import (
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	"mulberry/app/downstream/service"
	commonModel "mulberry/common/model"
	"mulberry/global"
)

func CreateRule(req *requests.CreateRule) error {
	srv := service.NewRule(global.DB, false)
	return srv.Create(&req.RuleInfo)
}

func UpdateRule(req *requests.UpdateRule) error {
	data := make(map[string]any)
	data["name"] = req.Name
	data["target_id"] = req.TargetId
	data["target_route"] = req.TargetRoute
	data["port"] = req.Port
	data["method"] = req.Method
	data["order"] = req.Order
	data["remark"] = req.Remark
	data["status"] = req.Status
	data["match_type"] = req.MatchType
	data["headers"] = req.Headers

	if req.IsWs {
		data["is_ws"] = 1
	} else {
		data["is_ws"] = 0
	}

	return service.NewRule(global.DB, false).UpdateByMap(req.ID, data)
}

func UpdateRuleStatus(id uint) error {
	srv := service.NewRule()
	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	return srv.UpdateByMap(id, map[string]any{"status": !info.Status})
}

func DeleteRule(id uint) error {
	return service.NewRule(global.DB, false).Delete(id)
}

func RuleList(condition *commonModel.PageQuery[*requests.QueryRule]) (*commonModel.ResPage[model.RuleInfo], error) {
	return service.NewRule(global.DB, false).ListRule(condition)
}

func GetRule(id uint) (*model.RuleInfo, error) {
	return service.NewRule(global.DB, false).GetById(id)
}
