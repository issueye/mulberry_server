package requests

type QueryCertDatas struct {
	Name   string `form:"name"`
	Status bool   `form:"status"`
}
