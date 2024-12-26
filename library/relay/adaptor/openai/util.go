package openai

import "demogogo/library/relay/model"

func ErrorWrapper(err error, code string, statusCode int) *model.ErrorWithStatusCode {
	Error := model.Error{
		Message: err.Error(),
		Type:    "error",
		Code:    code,
	}
	return &model.ErrorWithStatusCode{
		Error:      Error,
		StatusCode: statusCode,
	}
}
