package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	g.Meta `path:"/channel/test" tags:"channel" method:"get" summary:"渠道测试"`
	Id     int64 `v:"required" dc:"id" json:"id"`
	Model  int64 `v:"required" dc:"模型" json:"model"`
}

type TestRes struct {
	//g.Meta `mime:"text/html" example:"string"`
}
