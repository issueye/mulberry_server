package requests

import (
	"encoding/json"
	"mulberry/internal/app/downstream/model"
	commonModel "mulberry/internal/common/model"
)

type CreateTarget struct {
	model.TargetInfo
}

func NewCreateTarget() *CreateTarget {
	return &CreateTarget{
		TargetInfo: model.TargetInfo{},
	}
}

type UpdateTarget struct {
	model.TargetInfo
}

func NewUpdateTarget() *UpdateTarget {
	return &UpdateTarget{
		TargetInfo: model.TargetInfo{},
	}
}

func (req *CreateTarget) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryTarget struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryTarget() *commonModel.PageQuery[*QueryTarget] {
	return commonModel.NewPageQuery(0, 0, &QueryTarget{})
}
