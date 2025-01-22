package logic

import (
	"mulberry/internal/app/downstream/model"
	"mulberry/internal/app/downstream/requests"
	"mulberry/internal/app/downstream/service"
	commonModel "mulberry/internal/common/model"
	"mulberry/internal/global"
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
