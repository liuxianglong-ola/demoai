package relay

import (
	"bytes"
	"context"
	"demogogo/internal/dao"
	"demogogo/internal/model"
	"demogogo/internal/model/entity"
	"demogogo/internal/service"
	"demogogo/library/code"
	"demogogo/library/relay/adaptor"
	"demogogo/library/relay/meta"
	relaymodel "demogogo/library/relay/model"
	"demogogo/utility"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"strings"
)

//调用模型的服务

type (
	sRelay struct{}
)

func init() {
	service.RegisterRelay(New())
}

func New() service.IRelay {
	return &sRelay{}
}

func (s *sRelay) checkHeaderToken(ctx context.Context, i18n *model.I18n, key string, requestModel string) (tokenEntity *entity.Tokens, channelEntity *entity.Channels, err error) {
	if key == "" {
		//抛错
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayBuildRequestErr)
	}

	tokenEntity, err = dao.Tokens.GetRecordByKey(ctx, key)
	if err != nil {
		g.Log().Errorf(ctx, "sRelay.checkHeaderToken Tokens.GetRecordByKey err,key=%s,err=%v", key, err)

		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayTokenNoExistsErr)
	}
	if tokenEntity == nil {
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayTokenNoExistsErr)
	}
	if tokenEntity.Models == "" {
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayModelCannotUseErr)
	}
	modelSplit := strings.Split(tokenEntity.Models, ",")

	if requestModel != "" && !utility.InArray(requestModel, modelSplit) {
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayModelCannotUseErr)
	}

	//暂时不考虑黑名单，指定网段，指定group
	//去获取渠道id
	abilityEntity, err := dao.Abilities.GetRecordByModel(ctx, requestModel)
	if err != nil {
		g.Log().Errorf(ctx,
			"sRelay.checkHeaderToken Abilities.GetRecordByModel err,key=%s,model=%s,err=%v",
			key, requestModel, err)
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayBadModel)
	}
	if abilityEntity == nil {
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayBadModel)
	}
	channelEntity, err = dao.Channels.GetRecordById(ctx, abilityEntity.ChannelId)
	if err != nil {
		g.Log().Errorf(ctx,
			"sRelay.checkHeaderToken Channels.GetRecordById err,key=%s,channel=%d,err=%v",
			key, abilityEntity.ChannelId, err)
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayBadChannel)
	}
	if channelEntity == nil {
		return nil, nil, code.CodeError.ErrorNew(ctx, i18n, code.RelayBadChannel)
	}
	return
}

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
