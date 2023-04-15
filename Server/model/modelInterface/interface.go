package modelInterface

type PageInfo struct {
	Page     int
	PageSize int
}

// 分页接口
type Paging interface {
	GetInfoList(info PageInfo) (err error, list interface{}, total int)
}
