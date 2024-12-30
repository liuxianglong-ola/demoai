package doubao

import (
	"demogogo/library/relay/meta"
	"demogogo/library/relay/relaymode"
	"fmt"
)

func GetRequestURL(meta *meta.Meta) (string, error) {
	switch meta.Mode {
	case relaymode.ChatCompletions:
		return fmt.Sprintf("%s/api/v3/chat/completions", meta.BaseURL), nil
	case relaymode.Embeddings:
		return fmt.Sprintf("%s/api/v3/embeddings", meta.BaseURL), nil
	default:
	}
	return "", fmt.Errorf("unsupported relay mode %d for doubao", meta.Mode)
}
