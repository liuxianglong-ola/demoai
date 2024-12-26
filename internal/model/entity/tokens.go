// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Tokens is the golang structure for table tokens.
type Tokens struct {
	Id             int64  `json:"id"              orm:"id"              ` //
	UserId         int64  `json:"user_id"         orm:"user_id"         ` //
	Key            string `json:"key"             orm:"key"             ` //
	Status         int64  `json:"status"          orm:"status"          ` //
	Name           string `json:"name"            orm:"name"            ` //
	CreatedTime    int64  `json:"created_time"    orm:"created_time"    ` //
	AccessedTime   int64  `json:"accessed_time"   orm:"accessed_time"   ` //
	ExpiredTime    int64  `json:"expired_time"    orm:"expired_time"    ` //
	RemainQuota    int64  `json:"remain_quota"    orm:"remain_quota"    ` //
	UnlimitedQuota int    `json:"unlimited_quota" orm:"unlimited_quota" ` //
	UsedQuota      int64  `json:"used_quota"      orm:"used_quota"      ` //
	Models         string `json:"models"          orm:"models"          ` //
	Subnet         string `json:"subnet"          orm:"subnet"          ` //
}
