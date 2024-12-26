package channel

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"demogogo/app/http/api/channel/v1"
)

func (c *ControllerV1) Test(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
