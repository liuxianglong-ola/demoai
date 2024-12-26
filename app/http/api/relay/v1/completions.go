package v1

import "github.com/gogf/gf/v2/frame/g"

type CompletionsReq struct {
	g.Meta `path:"/completions" tags:"GPT" method:"post" summary:"completions"`
}

type CompletionsRes struct {
	//g.Meta `mime:"text/html" example:"string"`
}
