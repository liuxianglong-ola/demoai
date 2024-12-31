package relay

import (
	"demogogo/library/relay/adaptor"
	"demogogo/library/relay/adaptor/gemini"
	"demogogo/library/relay/adaptor/ollama"
	"demogogo/library/relay/adaptor/openai"
	"demogogo/library/relay/apitype"
)

func GetAdaptor(apiType int) adaptor.Adaptor {
	switch apiType {
	//case apitype.AIProxyLibrary:
	//	return &aiproxy.Adaptor{}
	//case apitype.Ali:
	//	return &ali.Adaptor{}
	//case apitype.Anthropic:
	//	return &anthropic.Adaptor{}
	//case apitype.AwsClaude:
	//	return &aws.Adaptor{}
	//case apitype.Baidu:
	//	return &baidu.Adaptor{}
	case apitype.Gemini:
		return &gemini.Adaptor{}
	case apitype.OpenAI:
		return &openai.Adaptor{}
	//case apitype.PaLM:
	//	return &palm.Adaptor{}
	//case apitype.Tencent:
	//	return &tencent.Adaptor{}
	//case apitype.Xunfei:
	//	return &xunfei.Adaptor{}
	//case apitype.Zhipu:
	//	return &zhipu.Adaptor{}
	case apitype.Ollama:
		return &ollama.Adaptor{}
		//case apitype.Coze:
		//	return &coze.Adaptor{}
		//case apitype.Cohere:
		//	return &cohere.Adaptor{}
		//case apitype.Cloudflare:
		//	return &cloudflare.Adaptor{}
		//case apitype.DeepL:
		//	return &deepl.Adaptor{}
		//case apitype.VertexAI:
		//	return &vertexai.Adaptor{}
		//case apitype.Proxy:
		//	return &proxy.Adaptor{}
	}
	return nil
}
