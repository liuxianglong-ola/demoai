package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EmbeddingsReq struct {
	g.Meta           `path:"/embeddings" tags:"GPT" method:"post" summary:"Embeddings"`
	Input            any      `v:"required" json:"input,omitempty"`
	Model            string   `v:"required" json:"model,omitempty"`
	EncodingFormat   string   `json:"encoding_format,omitempty"`
	Dimensions       int      `json:"dimensions,omitempty"`
	User             string   `json:"user,omitempty"`
	HeaderTokenKey   string   `v:"required" dc:"模型" json:"header_token_key"`
	Seed             float64  `json:"seed,omitempty"`
	Temperature      *float64 `json:"temperature,omitempty"`
	TopP             *float64 `json:"top_p,omitempty"`
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	PresencePenalty  *float64 `json:"presence_penalty,omitempty"`
}

type EmbeddingsRes struct {
	//g.Meta `mime:"text/html" example:"string"`
}
