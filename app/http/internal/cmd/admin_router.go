package cmd

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func adminRouter(group *ghttp.RouterGroup) {
	group.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Bind(
			//hello.NewV1(),
		)
	})
}
