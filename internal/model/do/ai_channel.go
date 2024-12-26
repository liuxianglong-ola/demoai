// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AiChannel is the golang structure of table ai_channel for DAO operations like Where/Data.
type AiChannel struct {
	g.Meta       `orm:"table:ai_channel, do:true"`
	Id           interface{} // 自增id
	Key          interface{} // 密钥
	Type         interface{} //
	Status       interface{} // 状态
	Name         interface{} // 名称
	CreatedTime  interface{} // 创建时间
	TestTime     interface{} // 测试时间
	ResponseTime interface{} // 测试响应时间
	BaseUrl      interface{} // 代理地址
	Other        interface{} //
	ModelStr     interface{} // 具体的模型
	Group        interface{} // 组
	Config       interface{} // 配置项
}
