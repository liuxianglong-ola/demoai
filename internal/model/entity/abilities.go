// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Abilities is the golang structure for table abilities.
type Abilities struct {
	Group     string `json:"group"      orm:"group"      ` //
	Model     string `json:"model"      orm:"model"      ` //
	ChannelId int64  `json:"channel_id" orm:"channel_id" ` //
	Enabled   int    `json:"enabled"    orm:"enabled"    ` //
	Priority  int64  `json:"priority"   orm:"priority"   ` //
}
