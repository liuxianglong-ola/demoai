package openai

import (
	//"demogogo/library/relay/adaptor/ai360"
	//"demogogo/library/relay/adaptor/baichuan"
	//"demogogo/library/relay/adaptor/deepseek"
	//"demogogo/library/relay/adaptor/doubao"
	//"demogogo/library/relay/adaptor/groq"
	//"demogogo/library/relay/adaptor/lingyiwanwu"
	//"demogogo/library/relay/adaptor/minimax"
	//"demogogo/library/relay/adaptor/mistral"
	//"demogogo/library/relay/adaptor/moonshot"
	//"demogogo/library/relay/adaptor/novita"
	//"demogogo/library/relay/adaptor/siliconflow"
	//"demogogo/library/relay/adaptor/stepfun"
	//"demogogo/library/relay/adaptor/togetherai"
	//"demogogo/library/relay/adaptor/xai"
	"demogogo/library/relay/channeltype"
)

var CompatibleChannels = []int{
	channeltype.Azure,
	channeltype.AI360,
	channeltype.Moonshot,
	channeltype.Baichuan,
	channeltype.Minimax,
	channeltype.Doubao,
	channeltype.Mistral,
	channeltype.Groq,
	channeltype.LingYiWanWu,
	channeltype.StepFun,
	channeltype.DeepSeek,
	channeltype.TogetherAI,
	channeltype.Novita,
	channeltype.SiliconFlow,
	channeltype.XAI,
}

func GetCompatibleChannelMeta(channelType int) (string, []string) {
	switch channelType {
	//@todo 用到时添加
	case channeltype.Azure:
		return "azure", ModelList
	//case channeltype.AI360:
	//	return "360", ai360.ModelList
	//case channeltype.Moonshot:
	//	return "moonshot", moonshot.ModelList
	//case channeltype.Baichuan:
	//	return "baichuan", baichuan.ModelList
	//case channeltype.Minimax:
	//	return "minimax", minimax.ModelList
	//case channeltype.Mistral:
	//	return "mistralai", mistral.ModelList
	//case channeltype.Groq:
	//	return "groq", groq.ModelList
	//case channeltype.LingYiWanWu:
	//	return "lingyiwanwu", lingyiwanwu.ModelList
	//case channeltype.StepFun:
	//	return "stepfun", stepfun.ModelList
	//case channeltype.DeepSeek:
	//	return "deepseek", deepseek.ModelList
	//case channeltype.TogetherAI:
	//	return "together.ai", togetherai.ModelList
	//case channeltype.Doubao:
	//	return "doubao", doubao.ModelList
	//case channeltype.Novita:
	//	return "novita", novita.ModelList
	//case channeltype.SiliconFlow:
	//	return "siliconflow", siliconflow.ModelList
	//case channeltype.XAI:
	//	return "xai", xai.ModelList
	default:
		return "openai", ModelList
	}
}
