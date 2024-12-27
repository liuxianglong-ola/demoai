package model

import (
	"demogogo/internal/model/entity"
	"demogogo/library/relay/channeltype"
	"demogogo/library/relay/meta"
	"demogogo/library/relay/model"
)

type ChatCompletions struct {
	Messages            []model.Message       `v:"required" dc:"消息" json:"messages"`
	Model               string                `v:"required" dc:"模型" json:"model"`
	Store               *bool                 `json:"store,omitempty"`
	Metadata            any                   `json:"metadata,omitempty"`
	FrequencyPenalty    *float64              `json:"frequency_penalty,omitempty"`
	LogitBias           any                   `json:"logit_bias,omitempty"`
	Logprobs            *bool                 `json:"logprobs,omitempty"`
	TopLogprobs         *int                  `json:"top_logprobs,omitempty"`
	MaxTokens           int                   `json:"max_tokens,omitempty"`
	MaxCompletionTokens *int                  `json:"max_completion_tokens,omitempty"`
	N                   int                   `json:"n,omitempty"`
	Modalities          []string              `json:"modalities,omitempty"`
	Prediction          any                   `json:"prediction,omitempty"`
	Audio               *model.Audio          `json:"audio,omitempty"`
	PresencePenalty     *float64              `json:"presence_penalty,omitempty"`
	ResponseFormat      *model.ResponseFormat `json:"response_format,omitempty"`
	Seed                float64               `json:"seed,omitempty"`
	ServiceTier         *string               `json:"service_tier,omitempty"`
	Stop                any                   `json:"stop,omitempty"`
	Stream              bool                  `json:"stream,omitempty"`
	StreamOptions       *model.StreamOptions  `json:"stream_options,omitempty"`
	Temperature         *float64              `json:"temperature,omitempty"`
	TopP                *float64              `json:"top_p,omitempty"`
	TopK                int                   `json:"top_k,omitempty"`
	Tools               []model.Tool          `json:"tools,omitempty"`
	ToolChoice          any                   `json:"tool_choice,omitempty"`
	ParallelTooCalls    *bool                 `json:"parallel_tool_calls,omitempty"`
	User                string                `json:"user,omitempty"`
	FunctionCall        any                   `json:"function_call,omitempty"`
	Functions           any                   `json:"functions,omitempty"`
}

func ChannelToMeta(token *entity.Tokens, channel *entity.Channels, mode int, reqModel string) *meta.Meta {
	meta := meta.Meta{
		Mode:            mode,
		ChannelType:     int(channel.Type),
		ChannelId:       int(channel.Id),
		TokenId:         int(token.Id),
		TokenName:       token.Name,
		UserId:          int(token.UserId),
		Group:           channel.Group,
		OriginModelName: reqModel,
		BaseURL:         channel.BaseUrl,
		APIKey:          channel.Key,
	}

	if meta.BaseURL == "" {
		meta.BaseURL = channeltype.ChannelBaseURLs[meta.ChannelType]
	}
	if meta.Method == "" {
		meta.Method = "POST"
	}
	meta.APIType = channeltype.ToAPIType(meta.ChannelType)
	return &meta
}
