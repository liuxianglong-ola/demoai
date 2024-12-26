package relay

import (
	"bytes"
	"context"
	v1 "demogogo/app/http/api/relay/v1"
	"demogogo/internal/model"
	"demogogo/internal/service"
	"demogogo/library/relay/adaptor"
	"demogogo/library/relay/meta"
	relaymodel "demogogo/library/relay/model"
	"demogogo/library/relay/relaymode"
	"demogogo/utility"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"io"
)

func (s *sRelay) getRequestBody(ctx context.Context, meta *meta.Meta, textRequest *relaymodel.GeneralOpenAIRequest, adaptor adaptor.Adaptor) (io.Reader, error) {
	//if meta.APIType == apitype.OpenAI && meta.OriginModelName == meta.ActualModelName && meta.ChannelType != channeltype.Baichuan {
	//	// no need to convert request for openai
	//	return c.Request.Body, nil
	//}

	// get request body
	var requestBody io.Reader
	convertedRequest, err := adaptor.ConvertRequest(ctx, meta.Mode, textRequest)
	if err != nil {
		g.Log().Debugf(ctx, "converted request failed: %s\n", err.Error())
		return nil, err
	}
	jsonData, err := json.Marshal(convertedRequest)
	if err != nil {
		g.Log().Debugf(ctx, "converted request json_marshal_failed: %s\n", err.Error())

		return nil, err
	}
	g.Log().Debugf(ctx, "converted request: \n%s", string(jsonData))

	requestBody = bytes.NewBuffer(jsonData)
	return requestBody, nil
}

func (s *sRelay) ChatCompletions(ctx context.Context, req *v1.ChatCompletionsReq) (err error) {
	i18n := service.BizCtx().Get(ctx).I18n
	token, channel, err := s.checkHeaderToken(ctx, i18n, req.HeaderTokenKey, req.Model)
	if err != nil {
		return err
	}
	mode := relaymode.ChatCompletions

	meta := model.ChannelToMeta(token, channel, mode, req.Model, req.Stream)

	textRequest := &relaymodel.GeneralOpenAIRequest{}
	utility.CopyFields(req, textRequest)
	return s.text(ctx, i18n, meta, textRequest)

}
