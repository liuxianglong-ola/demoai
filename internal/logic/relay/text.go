package relay

import (
	"context"
	"demogogo/internal/model"
	"demogogo/library/code"
	"demogogo/library/relay"
	"demogogo/library/relay/meta"
	relaymodel "demogogo/library/relay/model"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"net/http"
	"strconv"
)

func (s *sRelay) text(ctx context.Context, i18n *model.I18n, meta *meta.Meta, textRequest *relaymodel.GeneralOpenAIRequest) (err error) {

	adaptor := relay.GetAdaptor(meta.APIType)
	if adaptor == nil {
		//抛错
		g.Log().Errorf(ctx, "adaptor failed: %d", meta.APIType)
		return code.CodeError.ErrorNew(ctx, i18n, code.RelayBuildRequestErr)
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

	if resp != nil && resp.StatusCode != http.StatusOK {
		respErr := RelayErrorHandler(resp)
		if respErr != nil {
			g.Log().Errorf(ctx, "respErr is not nil: %+v", respErr)

			return code.CodeError.ErrorNew(ctx, i18n, code.CommonCustomizeError, respErr.Message)
		}
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

func RelayErrorHandler(resp *http.Response) (ErrorWithStatusCode *relaymodel.ErrorWithStatusCode) {
	if resp == nil {
		return &relaymodel.ErrorWithStatusCode{
			StatusCode: 500,
			Error: relaymodel.Error{
				Message: "resp is nil",
				Type:    "upstream_error",
				Code:    "bad_response",
			},
		}
	}
	ErrorWithStatusCode = &relaymodel.ErrorWithStatusCode{
		StatusCode: resp.StatusCode,
		Error: relaymodel.Error{
			Message: "",
			Type:    "upstream_error",
			Code:    "bad_response_status_code",
			Param:   strconv.Itoa(resp.StatusCode),
		},
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(responseBody))
	err = resp.Body.Close()
	if err != nil {
		return
	}
	var errResponse model.GeneralErrorResponse

	err = json.Unmarshal(responseBody, &errResponse)
	if err != nil {
		return
	}
	if errResponse.Error.Message != "" {
		// OpenAI format error, so we override the default one
		ErrorWithStatusCode.Error = errResponse.Error
	} else {
		ErrorWithStatusCode.Error.Message = errResponse.ToMessage()
	}
	if ErrorWithStatusCode.Error.Message == "" {
		ErrorWithStatusCode.Error.Message = fmt.Sprintf("bad response status code %d", resp.StatusCode)
	}
	return
}
