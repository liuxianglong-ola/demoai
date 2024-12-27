package relay

import (
	"context"
	"demogogo/internal/service"

	"demogogo/app/http/api/relay/v1"
)

func (c *ControllerV1) Embeddings(ctx context.Context, req *v1.EmbeddingsReq) (res *v1.EmbeddingsRes, err error) {
	err = service.Relay().Embeddings(ctx, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
