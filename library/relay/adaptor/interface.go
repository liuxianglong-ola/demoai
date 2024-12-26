package adaptor

import (
	"context"
	"demogogo/library/relay/meta"
	"demogogo/library/relay/model"
	"io"
	"net/http"
)

type Adaptor interface {
	Init(meta *meta.Meta)
	GetRequestURL(meta *meta.Meta) (string, error)
	SetupRequestHeader(c context.Context, req *http.Request, meta *meta.Meta) error
	ConvertRequest(c context.Context, relayMode int, request *model.GeneralOpenAIRequest) (any, error)
	ConvertImageRequest(request *model.ImageRequest) (any, error)
	DoRequest(c context.Context, meta *meta.Meta, requestBody io.Reader) (*http.Response, error)
	DoResponse(c context.Context, resp *http.Response, meta *meta.Meta) (usage *model.Usage, err *model.ErrorWithStatusCode)
	GetModelList() []string
	GetChannelName() string
}
