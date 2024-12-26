package main

import (
	"demogogo/app/http/internal/cmd"
	_ "demogogo/internal/boot"
	_ "demogogo/internal/logic"
	_ "demogogo/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
