package v1

import "github.com/gogf/gf/v2/frame/g"

type MessageReq struct {
	g.Meta       `path:"/chat-message" tags:"chat" method:"get" summary:"发送对话消息"`
	ResponseMode string
}

type MessageRes struct {
	//g.Meta `mime:"text/html" example:"string"`
}
