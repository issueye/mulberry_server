package requests

import commonModel "mulberry/common/model"

type QueryTraffic struct {
}

func NewQueryTraffic() *commonModel.PageQuery[*QueryTraffic] {
	return commonModel.NewPageQuery(0, 0, &QueryTraffic{})
}
