package logic

import (
	"mulberry/host/app/downstream/engine"
	"mulberry/host/app/downstream/model"
	"mulberry/host/app/downstream/requests"
	"mulberry/host/app/downstream/service"
	commonModel "mulberry/host/common/model"
	"mulberry/host/global"
)

func CreatePort(req *requests.CreatePort) error {
	srv := service.NewPort(global.DB, false)
	return srv.Create(&req.PortInfo)
}

func UpdatePort(req *requests.UpdatePort) error {
	return service.NewPort(global.DB, false).Update(req.ID, &req.PortInfo)
}

func DeletePort(id uint) error {
	return service.NewPort(global.DB, false).Delete(id)
}

func PortList(condition *commonModel.PageQuery[*requests.QueryPort]) (*commonModel.ResPage[model.PortInfo], error) {
	return service.NewPort(global.DB, false).ListPort(condition)
}

func GetPort(id uint) (*model.PortInfo, error) {
	return service.NewPort(global.DB, false).GetById(id)
}

func Reload(port uint) error {
	info, err := service.NewPort().GetByField("port", port)
	if err != nil {
		return err
	}

	engine.PortChan <- &engine.Port{
		PortInfo: *info,
		Action:   engine.AT_RELOAD,
	}

	return nil
}

func Start(port uint) error {
	srv := service.NewPort()
	info, err := srv.GetByField("port", port)
	if err != nil {
		return err
	}

	engine.PortChan <- &engine.Port{
		PortInfo: *info,
		Action:   engine.AT_START,
	}

	// 修改状态
	return srv.UpdateByField("port", port, map[string]any{"status": 1})
}

func Stop(port uint) error {
	srv := service.NewPort()

	info, err := srv.GetByField("port", port)
	if err != nil {
		return err
	}

	engine.PortChan <- &engine.Port{
		PortInfo: *info,
		Action:   engine.AT_STOP,
	}

	// 修改状态
	return srv.UpdateByField("port", port, map[string]any{"status": 0})
}

func ModifyUseGZ(port uint) error {
	srv := service.NewPort()

	info, err := srv.GetByField("port", port)
	if err != nil {
		return err
	}

	// 修改状态
	return srv.UpdateByField("port", port, map[string]any{"use_gzip": !info.UseGzip})
}
