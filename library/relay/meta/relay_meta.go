package meta

import (
	"demogogo/library/relay/channeltype"
	"demogogo/library/relay/common/ctxkey"
	"demogogo/library/relay/model"
	"demogogo/library/relay/relaymode"
	"github.com/gin-gonic/gin"
	"strings"
)

type Meta struct {
	Mode         int
	ChannelType  int
	ChannelId    int
	TokenId      int
	TokenName    string
	UserId       int
	Group        string
	ModelMapping map[string]string
	// BaseURL is the proxy url set in the channel config
	BaseURL  string
	APIKey   string
	APIType  int
	Config   model.ChannelConfig
	IsStream bool
	// OriginModelName is the model name from the raw user request
	OriginModelName string
	// ActualModelName is the model name after mapping
	ActualModelName string
	RequestURLPath  string
	PromptTokens    int // only for DoResponse
	SystemPrompt    string
	ContentType     string
	Accept          string
	Method          string
}

func GetByContext(c *gin.Context) *Meta {
	meta := Meta{
		Mode:            relaymode.GetByPath(c.Request.URL.Path),
		ChannelType:     c.GetInt(ctxkey.Channel),
		ChannelId:       c.GetInt(ctxkey.ChannelId),
		TokenId:         c.GetInt(ctxkey.TokenId),
		TokenName:       c.GetString(ctxkey.TokenName),
		UserId:          c.GetInt(ctxkey.Id),
		Group:           c.GetString(ctxkey.Group),
		ModelMapping:    c.GetStringMapString(ctxkey.ModelMapping),
		OriginModelName: c.GetString(ctxkey.RequestModel),
		BaseURL:         c.GetString(ctxkey.BaseURL),
		APIKey:          strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer "),
		RequestURLPath:  c.Request.URL.String(),
		SystemPrompt:    c.GetString(ctxkey.SystemPrompt),
	}
	cfg, ok := c.Get(ctxkey.Config)
	if ok {
		meta.Config = cfg.(model.ChannelConfig)
	}
	if meta.BaseURL == "" {
		meta.BaseURL = channeltype.ChannelBaseURLs[meta.ChannelType]
	}
	meta.APIType = channeltype.ToAPIType(meta.ChannelType)
	return &meta
}