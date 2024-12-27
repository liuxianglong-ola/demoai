// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package relay

import (
	"context"

	"demogogo/app/http/api/relay/v1"
)

type IRelayV1 interface {
	ChatCompletions(ctx context.Context, req *v1.ChatCompletionsReq) (res *v1.ChatCompletionsRes, err error)
	Completions(ctx context.Context, req *v1.CompletionsReq) (res *v1.CompletionsRes, err error)
	Embeddings(ctx context.Context, req *v1.EmbeddingsReq) (res *v1.EmbeddingsRes, err error)
}
