package middleware

import (
	"demogogo/library/code"
	"demogogo/protobuf/pb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"strconv"
	"strings"
)

func (s *sMiddleware) AiTokenCheck(r *ghttp.Request) {
	key := r.GetHeader("Authorization")
	key = strings.TrimPrefix(key, "Bearer ")
	key = strings.TrimPrefix(key, "sk-")
	//检查token存在不
	if key == "" {
		//抛错
		err := code.CodeError.New(r.GetCtx(), code.RelayTokenNoExistsErr)
		codeErr := gerror.Code(err)

		s.buildAiErrorResponse(r, codeErr.Code(), err.Error())
		return
	}
	r.SetQuery("header_token_key", key)

	r.Middleware.Next()
}
func (s *sMiddleware) AiResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg string
		err = r.GetError()
		//res     = r.GetHandlerResponse()
		codeErr = gerror.Code(err)
	)
	if err != nil {
		if codeErr == gcode.CodeNil {
			codeErr = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				codeErr = gcode.CodeNotFound
			case http.StatusForbidden:
				codeErr = gcode.CodeNotAuthorized
			default:
				codeErr = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(codeErr, msg)
			r.SetError(err)
		} else {
			return
		}
	}

	r.Response.Flush()
	s.buildAiErrorResponse(r, codeErr.Code(), err.Error())
	r.ExitAll()
	return
}

func (s *sMiddleware) buildAiErrorResponse(r *ghttp.Request, code int, message string) {
	r.Response.Status = http.StatusInternalServerError
	Error := &pb.RelayErrDetailData{
		Message: message,
		Type:    "error",
		Code:    strconv.Itoa(code),
	}

	r.Response.WriteJson(&pb.RelayErrData{Error: Error})
}
