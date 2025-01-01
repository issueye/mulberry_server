package logic

import (
	"fmt"
	"mulberry/app/downstream/engine"
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	"mulberry/app/downstream/service"
	commonModel "mulberry/common/model"
	"mulberry/global"
)

func CreatePort(req *requests.CreatePort) error {
	srv := service.NewPort(global.DB, false)
	return srv.Create(&req.PortInfo)
}

func UpdatePort(req *requests.UpdatePort) error {
	return service.NewPort(global.DB, false).Update(req.ID, &req.PortInfo)
}

func DeletePort(id uint) error {
	var (
		err   error
		pInfo *model.PortInfo
	)

	db := global.DB.Begin()
	pSrv := service.NewPort(db, true)
	defer func() {
		if err != nil {
			err := pSrv.Rollback()
			if err != nil {
				global.Logger.Sugar().Errorf("数据回滚失败 %s", err.Error())
			}

			return
		}

		err := pSrv.Commit()
		if err != nil {
			global.Logger.Sugar().Errorf("提交事务失败 %s", err.Error())
		}
	}()

	pInfo, err = pSrv.GetById(id)
	if err != nil {
		return err
	}

	if pInfo.Status {
		err = fmt.Errorf("端口[%d]正在监听中...不能删除", pInfo.Port)
		return err
	}

	err = pSrv.Delete(id)
	if err != nil {
		return err
	}

	pageSrv := service.NewPage(pSrv.TX, true)
	err = pageSrv.DeleteByFields(map[string]any{"port": pInfo.Port})
	if err != nil {
		return err
	}

	ruleSrv := service.NewRule(pSrv.TX, true)
	err = ruleSrv.DeleteByFields(map[string]any{"port": pInfo.Port})
	return err
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
