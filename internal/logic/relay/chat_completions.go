package relay

import (
	"context"
	v1 "demogogo/app/http/api/relay/v1"
	"demogogo/internal/model"
	"demogogo/internal/service"
	relaymodel "demogogo/library/relay/model"
	"demogogo/library/relay/relaymode"
	"demogogo/utility"
)

func (s *sRelay) ChatCompletions(ctx context.Context, req *v1.ChatCompletionsReq) (err error) {
	i18n := service.BizCtx().Get(ctx).I18n
	token, channel, err := s.checkHeaderToken(ctx, i18n, req.HeaderTokenKey, req.Model)
	if err != nil {
		return err
	}
	mode := relaymode.ChatCompletions

	meta := model.ChannelToMeta(ctx, token, channel, mode, req.Model)
	meta.IsStream = req.Stream
	req.Model = meta.ActualModelName
	textRequest := &relaymodel.GeneralOpenAIRequest{}
	utility.CopyFields(req, textRequest)
	return s.text(ctx, i18n, meta, textRequest)
}

func (s *sRelay) Embeddings(ctx context.Context, req *v1.EmbeddingsReq) (err error) {
	i18n := service.BizCtx().Get(ctx).I18n
	token, channel, err := s.checkHeaderToken(ctx, i18n, req.HeaderTokenKey, req.Model)
	if err != nil {
		return err
	}
	mode := relaymode.Embeddings

	meta := model.ChannelToMeta(ctx, token, channel, mode, req.Model)
	req.Model = meta.ActualModelName
	textRequest := &relaymodel.GeneralOpenAIRequest{}
	utility.CopyFields(req, textRequest)
	return s.text(ctx, i18n, meta, textRequest)
}
