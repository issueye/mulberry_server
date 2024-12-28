package logic

import (
	"mulberry/host/app/downstream/model"
	"mulberry/host/app/downstream/requests"
	"mulberry/host/app/downstream/service"
	commonModel "mulberry/host/common/model"
	"mulberry/host/global"
)

func CreateTarget(req *requests.CreateTarget) error {
	srv := service.NewTarget(global.DB, false)
	return srv.Create(&req.TargetInfo)
}

func UpdateTarget(req *requests.UpdateTarget) error {
	return service.NewTarget(global.DB, false).Update(req.ID, &req.TargetInfo)
}

func DeleteTarget(id uint) error {
	return service.NewTarget(global.DB, false).Delete(id)
}

func TargetList(condition *commonModel.PageQuery[*requests.QueryTarget]) (*commonModel.ResPage[model.TargetInfo], error) {
	return service.NewTarget(global.DB, false).ListTarget(condition)
}

func GetTarget(id uint) (*model.TargetInfo, error) {
	return service.NewTarget(global.DB, false).GetById(id)
}

func ModifyStatusTarget(id uint) error {
	srv := service.NewTarget()
	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	return srv.UpdateByMap(id, map[string]any{"status": !info.Status})
}
