package relay

import (
	"context"
	"demogogo/internal/consts"
	"demogogo/internal/logic/bizctx"
	"demogogo/internal/model"
	"demogogo/internal/service"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func Test_logic_relay_view(t *testing.T) {
	ctx := context.Background()
	service.RegisterBizCtx(bizctx.New())

	customCtx := &model.Context{
		I18n: model.NewI18n(),
		Data: make(g.Map),
		User: &model.ContextUser{},
	}
	//
	ctx = context.WithValue(ctx, consts.ContextKey, customCtx)

	s := &sRelay{}
	err := s.ChatCompletions(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("suc")
	}
}
