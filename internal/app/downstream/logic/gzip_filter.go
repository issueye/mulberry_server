package logic

import (
	"mulberry/internal/app/downstream/model"
	"mulberry/internal/app/downstream/requests"
	"mulberry/internal/app/downstream/service"
	commonModel "mulberry/internal/common/model"
)

func CreateGzipFilter(req *requests.CreateGzipFilter) error {
	srv := service.NewGzipFilter()
	return srv.Create(&req.GzipFilterInfo)
}

func UpdateGzipFilter(req *requests.UpdateGzipFilter) error {
	return service.NewGzipFilter().Update(req.ID, &req.GzipFilterInfo)
}

func DeleteGzipFilter(id uint) error {
	return service.NewGzipFilter().Delete(id)
}

func GzipFilterList(condition *commonModel.PageQuery[*requests.QueryGzipFilter]) (*commonModel.ResPage[model.GzipFilterInfo], error) {
	return service.NewGzipFilter().ListGzipFilter(condition)
}

func GetGzipFilter(id uint) (*model.GzipFilterInfo, error) {
	return service.NewGzipFilter().GetById(id)
}

func ModifyStatusGzipFilter(id uint) error {
	srv := service.NewGzipFilter()
	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	return srv.UpdateByMap(id, map[string]any{"status": !info.Status})
}
