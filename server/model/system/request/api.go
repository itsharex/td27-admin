package request

import "server/model/common/request"

type apiStruct struct {
	Path        string `json:"path"`
	Group       string `json:"group"`
	Method      string `json:"method" validate:"oneof=GET POST DELETE PUT"`
	Description string `json:"description"`
}

// ApiSearchParams api分页条件查询及排序结构体
type ApiSearchParams struct {
	apiStruct
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
