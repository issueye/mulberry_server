package logic

import (
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	"mulberry/app/downstream/service"
	commonModel "mulberry/common/model"
	"mulberry/global"
	"mulberry/pkg/utils"
	"path/filepath"
	"strconv"
)

func CreatePage(req *requests.CreatePage) error {
	srv := service.NewPage(global.DB, false)
	return srv.Create(&req.PageInfo)
}

func UpdatePage(req *requests.UpdatePage) error {
	data := make(map[string]any)
	data["name"] = req.Name
	data["title"] = req.Title
	data["version"] = req.Version
	data["port"] = req.Port
	data["product_code"] = req.ProductCode
	data["use_version_route"] = req.UseVersionRoute
	data["status"] = req.Status
	data["remark"] = req.Remark

	return service.NewPage(global.DB, false).UpdateByMap(req.ID, data)
}

func DeletePage(id uint) error {
	return service.NewPage(global.DB, false).Delete(id)
}

func PageList(condition *commonModel.PageQuery[*requests.QueryPage]) (*commonModel.ResPage[model.PageInfo], error) {
	return service.NewPage(global.DB, false).ListPage(condition)
}

func GetPage(id uint) (*model.PageInfo, error) {
	return service.NewPage(global.DB, false).GetById(id)
}

func SaveVersionPage(req *requests.SaveVersionPage) error {
	pageInfo, err := service.NewPage().GetByMap(map[string]any{
		"product_code": req.ProductCode,
		"port":         req.Port,
	})

	if err != nil {
		return err
	}

	// 读取页面文件
	path := filepath.Join(global.ROOT_PATH, req.Path)
	// 添加端口号作为路径，进行文件分离
	targetPath := filepath.Join(global.ROOT_PATH, "pages", strconv.Itoa(int(pageInfo.Port)), pageInfo.Name, req.Version)
	err = utils.Unzip(path, targetPath)
	if err != nil {
		return err
	}

	data := &model.PageVersionInfo{}
	data.PagePath = targetPath
	data.Version = req.Version
	data.Mark = ""
	data.PageId = pageInfo.ID

	err = service.NewPageVersion().Create(data)
	if err != nil {
		return err
	}

	// 反写版本号到页面信息
	pageInfo.Version = req.Version
	return service.NewPage().Update(pageInfo.ID, pageInfo)
}

func GetVersionList(pageId uint) ([]*model.PageVersionInfo, error) {
	return service.NewPageVersion().GetDatasByMap(map[string]any{
		"page_id": pageId,
	})
}

func UpdatePageStatus(id uint) error {
	srv := service.NewPage()
	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	return srv.UpdateByMap(id, map[string]any{"status": !info.Status})
}
