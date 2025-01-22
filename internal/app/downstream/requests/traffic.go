package requests

import commonModel "mulberry/internal/common/model"

type QueryTraffic struct {
}

func NewQueryTraffic() *commonModel.PageQuery[*QueryTraffic] {
	return commonModel.NewPageQuery(0, 0, &QueryTraffic{})
}
