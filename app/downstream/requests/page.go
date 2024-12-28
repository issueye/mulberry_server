package requests

import (
	"encoding/json"
	"mulberry/host/app/downstream/model"
	commonModel "mulberry/host/common/model"
)

type CreatePage struct {
	model.PageInfo
}

func NewCreatePage() *CreatePage {
	return &CreatePage{
		PageInfo: model.PageInfo{},
	}
}

type UpdatePage struct {
	model.PageInfo
}

func NewUpdatePage() *UpdatePage {
	return &UpdatePage{
		PageInfo: model.PageInfo{},
	}
}

func (req *CreatePage) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryPage struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
	Port     uint   `json:"port" form:"port"`         // 端口
}

func NewQueryPage() *commonModel.PageQuery[*QueryPage] {
	return commonModel.NewPageQuery(0, 0, &QueryPage{})
}

type SaveVersionPage struct {
	Version     string `binding:"required" label:"版本" json:"version"`        // 版本
	Port        uint   `binding:"required" label:"端口号" json:"port"`          // 端口信息编码
	ProductCode string `binding:"required" label:"产品代码" json:"product_code"` // 产品代码
	Path        string `binding:"required" label:"页面路径" json:"path"`         // 页面路径
}
