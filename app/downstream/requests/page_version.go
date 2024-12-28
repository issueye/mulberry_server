package requests

import (
	"encoding/json"
	"mulberry/host/app/downstream/model"
	commonModel "mulberry/host/common/model"
)

type CreatePageVersion struct {
	model.PageVersionInfo
}

func NewCreatePageVersion() *CreatePageVersion {
	return &CreatePageVersion{
		PageVersionInfo: model.PageVersionInfo{},
	}
}

type UpdatePageVersion struct {
	model.PageVersionInfo
}

func NewUpdatePageVersion() *UpdatePageVersion {
	return &UpdatePageVersion{
		PageVersionInfo: model.PageVersionInfo{},
	}
}

func (req *CreatePageVersion) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryPageVersion struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
	Port     uint   `json:"port" form:"port"`         // 端口
}

func NewQueryPageVersion() *commonModel.PageQuery[*QueryPageVersion] {
	return commonModel.NewPageQuery(0, 0, &QueryPageVersion{})
}
