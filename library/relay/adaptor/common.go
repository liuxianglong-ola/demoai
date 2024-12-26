package adaptor

import (
	"context"
	"demogogo/library/relay/meta"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"net/http"
)

func SetupCommonRequestHeader(c context.Context, req *http.Request, meta *meta.Meta) {
	if meta.ContentType != "" {
		req.Header.Set("Content-Type", meta.ContentType)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
	if meta.Accept != "" {
		req.Header.Set("Accept", meta.Accept)
	}
	if meta.IsStream && meta.Accept == "" {
		req.Header.Set("Accept", "text/event-stream")
	}
	//req.Header.Set("Content-Type", c.Request.Header.Get("Content-Type"))
	//req.Header.Set("Accept", c.Request.Header.Get("Accept"))
	//if meta.IsStream && c.Request.Header.Get("Accept") == "" {
	//	req.Header.Set("Accept", "text/event-stream")
	//}
}

func DoRequestHelper(a Adaptor, c context.Context, meta *meta.Meta, requestBody io.Reader) (*http.Response, error) {
	fullRequestURL, err := a.GetRequestURL(meta)
	if err != nil {
		return nil, fmt.Errorf("get request url failed: %w", err)
	}
	req, err := http.NewRequest(meta.Method, fullRequestURL, requestBody)

	if err != nil {
		return nil, fmt.Errorf("new request failed: %w", err)
	}
	err = a.SetupRequestHeader(c, req, meta)
	if err != nil {
		return nil, fmt.Errorf("setup request header failed: %w", err)
	}
	resp, err := DoRequest(c, req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	return resp, nil
}

func DoRequest(c context.Context, req *http.Request) (*http.Response, error) {
	resp, err := g.Client().Do(req)
	//if err != nil {
	//	return nil, err
	//}
	//resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("resp is nil")
	}
	_ = req.Body.Close()
	//_ = c.Request.Body.Close()

	return resp, nil
}
