package channel

import (
	"context"
	v1 "demogogo/app/http/api/channel/v1"
	relaymodel "demogogo/library/relay/model"
)

func buildTestRequest(modelName string) *relaymodel.GeneralOpenAIRequest {
	if modelName == "" {
		modelName = "gpt-3.5-turbo"
	}

	testRequet := &relaymodel.GeneralOpenAIRequest{
		MaxTokens: 2,
		Model:     modelName,
	}
	testMessage := relaymodel.Message{
		Role:    "user",
		Content: "hi",
	}
	testRequet.Messages = append(testRequet.Messages, testMessage)
	return testRequet
}

func (s *sChannel) Test(ctx context.Context, req *v1.TestReq) (err error, openaiErr *relaymodel.Error) {
	//查看数据在不在

	//meta := meta.GetByContext(c)
	//adaptor := relay.GetAdaptor(req.Type)
	//if adaptor == nil {
	//	return fmt.Errorf("invalid api type: %d, adaptor is nil", req.Type), nil
	//}
	//testRequet := buildTestRequest(req.Model)
	//
	//adaptor.Init(meta)
	//
	//convertedRequest, err := adaptor.ConvertRequest(c, relaymode.ChatCompletions, request)
	//
	//if err != nil {
	//	return err, nil
	//}
	//jsonData, err := json.Marshal(convertedRequest)
	//if err != nil {
	//	return err, nil
	//}
	//logger.SysLog(string(jsonData))
	//requestBody := bytes.NewBuffer(jsonData)
	//c.Request.Body = io.NopCloser(requestBody)
	//resp, err := adaptor.DoRequest(c, meta, requestBody)
	//if err != nil {
	//	return err, nil
	//}
	//if resp != nil && resp.StatusCode != http.StatusOK {
	//	//err := controller.RelayErrorHandler(resp)
	//	//return fmt.Errorf("status code %d: %s", resp.StatusCode, err.Error.Message), &err.Error
	//}
	//usage, respErr := adaptor.DoResponse(c, resp, meta)
	//if respErr != nil {
	//	return fmt.Errorf("%s", respErr.Error.Message), &respErr.Error
	//}
	//if usage == nil {
	//	return errors.New("usage is nil"), nil
	//}
	//result := w.Result()
	//// print result.Body
	//respBody, err := io.ReadAll(result.Body)
	//if err != nil {
	//	return err, nil
	//}
	//logger.SysLog(fmt.Sprintf("testing channel #%d, response: \n%s", channel.Id, string(respBody)))
	return nil, nil
}
