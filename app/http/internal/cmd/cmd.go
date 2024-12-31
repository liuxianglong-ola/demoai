package cmd

import (
	"context"
	"demogogo/internal/service"
	"demogogo/library/relay/adaptor/openai"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			openai.InitTokenEncoders()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
				)
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				adminRouter(group)
				apiRouter(group)
			})
			s.Run()
			return nil
		},
	}
)
