package relay

import (
	"context"
	"demogogo/app/http/api/relay/v1"
	"demogogo/internal/service"
)

func (c *ControllerV1) ChatCompletions(ctx context.Context, req *v1.ChatCompletionsReq) (res *v1.ChatCompletionsRes, err error) {

	err = service.Relay().ChatCompletions(ctx, req)
	if err != nil {
		return nil, err
	}
	return nil, nil

}
