// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demogogo/app/http/api/relay/v1"
)

type (
	IRelay interface {
		ChatCompletions(ctx context.Context, req *v1.ChatCompletionsReq) (err error)
		Embeddings(ctx context.Context, req *v1.EmbeddingsReq) (err error)
	}
)

var (
	localRelay IRelay
)

func Relay() IRelay {
	if localRelay == nil {
		panic("implement not found for interface IRelay, forgot register?")
	}
	return localRelay
}

func RegisterRelay(i IRelay) {
	localRelay = i
}
