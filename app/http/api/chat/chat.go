// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chat

import (
	"context"

	"demogogo/app/http/api/chat/v1"
)

type IChatV1 interface {
	Message(ctx context.Context, req *v1.MessageReq) (res *v1.MessageRes, err error)
}
