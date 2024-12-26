package v1

import (
	"demogogo/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ChatCompletionsReq struct {
	g.Meta `path:"/chat/completions" tags:"GPT" method:"post" summary:"chat-completions"`
	model.ChatCompletions
	HeaderTokenKey string `v:"required" dc:"模型" json:"header_token_key"`
}

type ChatCompletionsRes struct {
	//g.Meta `mime:"text/html" example:"string"`
}
