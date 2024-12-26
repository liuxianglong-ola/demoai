package relay

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"demogogo/app/http/api/relay/v1"
)

func (c *ControllerV1) Completions(ctx context.Context, req *v1.CompletionsReq) (res *v1.CompletionsRes, err error) {
	//replayMode := relaymode.Completions
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
