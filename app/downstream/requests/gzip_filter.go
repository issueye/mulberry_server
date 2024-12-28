package requests

import (
	"encoding/json"
	"mulberry/host/app/downstream/model"
	commonModel "mulberry/host/common/model"
)

type CreateGzipFilter struct {
	model.GzipFilterInfo
}

func NewCreateGzipFilter() *CreateGzipFilter {
	return &CreateGzipFilter{
		GzipFilterInfo: model.GzipFilterInfo{},
	}
}

type UpdateGzipFilter struct {
	model.GzipFilterInfo
}

func NewUpdateGzipFilter() *UpdateGzipFilter {
	return &UpdateGzipFilter{
		GzipFilterInfo: model.GzipFilterInfo{},
	}
}

func (req *CreateGzipFilter) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryGzipFilter struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryGzipFilter() *commonModel.PageQuery[*QueryGzipFilter] {
	return commonModel.NewPageQuery(0, 0, &QueryGzipFilter{})
}
