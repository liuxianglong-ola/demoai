// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AiChannel is the golang structure for table ai_channel.
type AiChannel struct {
	Id           int64  `json:"id"            orm:"id"            ` // 自增id
	Key          string `json:"key"           orm:"key"           ` // 密钥
	Type         int    `json:"type"          orm:"type"          ` //
	Status       int    `json:"status"        orm:"status"        ` // 状态
	Name         string `json:"name"          orm:"name"          ` // 名称
	CreatedTime  int64  `json:"created_time"  orm:"created_time"  ` // 创建时间
	TestTime     int64  `json:"test_time"     orm:"test_time"     ` // 测试时间
	ResponseTime int64  `json:"response_time" orm:"response_time" ` // 测试响应时间
	BaseUrl      string `json:"base_url"      orm:"base_url"      ` // 代理地址
	Other        string `json:"other"         orm:"other"         ` //
	ModelStr     string `json:"model_str"     orm:"model_str"     ` // 具体的模型
	Group        string `json:"group"         orm:"group"         ` // 组
	Config       string `json:"config"        orm:"config"        ` // 配置项
}
