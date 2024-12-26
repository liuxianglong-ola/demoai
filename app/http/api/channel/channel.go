// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package channel

import (
	"context"

	"demogogo/app/http/api/channel/v1"
)

type IChannelV1 interface {
	Test(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error)
}
