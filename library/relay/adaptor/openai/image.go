package openai

import (
	"context"
	"demogogo/library/relay/model"
	"net/http"
)

func ImageHandler(c context.Context, resp *http.Response) (*model.ErrorWithStatusCode, *model.Usage) {
	//@todo 临时注掉，到图片这块再打开
	//var imageResponse ImageResponse
	//responseBody, err := io.ReadAll(resp.Body)
	//
	//if err != nil {
	//	return ErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError), nil
	//}
	//err = resp.Body.Close()
	//if err != nil {
	//	return ErrorWrapper(err, "close_response_body_failed", http.StatusInternalServerError), nil
	//}
	//err = json.Unmarshal(responseBody, &imageResponse)
	//if err != nil {
	//	return ErrorWrapper(err, "unmarshal_response_body_failed", http.StatusInternalServerError), nil
	//}
	//
	//resp.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	//
	//for k, v := range resp.Header {
	//	c.Writer.Header().Set(k, v[0])
	//}
	//c.Writer.WriteHeader(resp.StatusCode)
	//
	//_, err = io.Copy(c.Writer, resp.Body)
	//if err != nil {
	//	return ErrorWrapper(err, "copy_response_body_failed", http.StatusInternalServerError), nil
	//}
	//err = resp.Body.Close()
	//if err != nil {
	//	return ErrorWrapper(err, "close_response_body_failed", http.StatusInternalServerError), nil
	//}
	return nil, nil
}
