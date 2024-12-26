package render

import (
	"context"
	"demogogo/library/relay/common"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

func StringData(ctx context.Context, str string) {
	str = strings.TrimPrefix(str, "data: ")
	str = strings.TrimSuffix(str, "\r")
	c := g.RequestFromCtx(ctx)
	event := common.CustomEvent{Data: "data: " + str}
	event.Render(c.Response.Writer)

	c.Response.Flush()
}

func ObjectData(c context.Context, object interface{}) error {
	jsonData, err := json.Marshal(object)
	if err != nil {
		return fmt.Errorf("error marshalling object: %w", err)
	}
	StringData(c, string(jsonData))
	return nil
}

func Done(c context.Context) {
	StringData(c, "[DONE]")
}
