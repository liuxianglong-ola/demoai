package cmd

import (
	"demogogo/app/http/internal/controller/hello"
	"demogogo/app/http/internal/controller/relay"
	"demogogo/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

func apiRouter(group *ghttp.RouterGroup) {
	group.Middleware(
		service.Middleware().AiTokenCheck,
		service.Middleware().AiResponse,
	)
	group.Bind(
		hello.NewV1(),
		relay.NewV1(),
	)
}
