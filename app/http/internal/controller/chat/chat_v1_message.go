package chat

import (
	"context"
	"demogogo/internal/service"

	"demogogo/app/http/api/chat/v1"
)

func (c *ControllerV1) Message(ctx context.Context, req *v1.MessageReq) (res *v1.MessageRes, err error) {
	service.Relay().ChatCompletions(ctx)
	return nil, nil
}
