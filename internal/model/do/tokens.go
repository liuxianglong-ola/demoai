// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Tokens is the golang structure of table tokens for DAO operations like Where/Data.
type Tokens struct {
	g.Meta         `orm:"table:tokens, do:true"`
	Id             interface{} //
	UserId         interface{} //
	Key            interface{} //
	Status         interface{} //
	Name           interface{} //
	CreatedTime    interface{} //
	AccessedTime   interface{} //
	ExpiredTime    interface{} //
	RemainQuota    interface{} //
	UnlimitedQuota interface{} //
	UsedQuota      interface{} //
	Models         interface{} //
	Subnet         interface{} //
}
