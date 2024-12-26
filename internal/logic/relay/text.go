package relay

import (
	"context"
	"demogogo/internal/model"
	"demogogo/library/code"
	"demogogo/library/relay"
	"demogogo/library/relay/meta"
	relaymodel "demogogo/library/relay/model"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sRelay) text(ctx context.Context, i18n *model.I18n, meta *meta.Meta, textRequest *relaymodel.GeneralOpenAIRequest) (err error) {

	adaptor := relay.GetAdaptor(meta.APIType)
	if adaptor == nil {
		//抛错
		return nil
		//openai.ErrorWrapper(fmt.Errorf("invalid api type: %d", meta.APIType), "invalid_api_type", http.StatusBadRequest)
	}
	adaptor.Init(meta)

	// get request body
	requestBody, err := s.getRequestBody(ctx, meta, textRequest, adaptor)
	if err != nil {
		g.Log().Errorf(ctx, "convert failed: %s", err.Error())
		return code.CodeError.ErrorNew(ctx, i18n, code.RelayBuildRequestErr)
	}

	// do request
	resp, err := adaptor.DoRequest(ctx, meta, requestBody)
	if err != nil {
		g.Log().Errorf(ctx, "DoRequest failed: %s", err.Error())
		return code.CodeError.ErrorNew(ctx, i18n, code.RelayRequestErr)
	}

	// do response
	usage, respErr := adaptor.DoResponse(ctx, resp, meta)
	if respErr != nil {
		g.Log().Errorf(ctx, "respErr is not nil: %+v", respErr)

		return code.CodeError.ErrorNew(ctx, i18n, code.CommonCustomizeError, respErr.Message)
	}
	//用掉的token
	fmt.Println(usage)
	return nil
}
