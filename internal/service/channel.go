// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demogogo/app/http/api/channel/v1"
	relaymodel "demogogo/library/relay/model"
)

type (
	IChannel interface {
		Test(ctx context.Context, req *v1.TestReq) (err error, openaiErr *relaymodel.Error)
	}
)

var (
	localChannel IChannel
)

func Channel() IChannel {
	if localChannel == nil {
		panic("implement not found for interface IChannel, forgot register?")
	}
	return localChannel
}

func RegisterChannel(i IChannel) {
	localChannel = i
}
