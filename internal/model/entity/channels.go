// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Channels is the golang structure for table channels.
type Channels struct {
	Id                 int64   `json:"id"                   orm:"id"                   ` //
	Type               int64   `json:"type"                 orm:"type"                 ` //
	Key                string  `json:"key"                  orm:"key"                  ` //
	Status             int64   `json:"status"               orm:"status"               ` //
	Name               string  `json:"name"                 orm:"name"                 ` //
	Weight             uint64  `json:"weight"               orm:"weight"               ` //
	CreatedTime        int64   `json:"created_time"         orm:"created_time"         ` //
	TestTime           int64   `json:"test_time"            orm:"test_time"            ` //
	ResponseTime       int64   `json:"response_time"        orm:"response_time"        ` //
	BaseUrl            string  `json:"base_url"             orm:"base_url"             ` //
	Other              string  `json:"other"                orm:"other"                ` //
	Balance            float64 `json:"balance"              orm:"balance"              ` //
	BalanceUpdatedTime int64   `json:"balance_updated_time" orm:"balance_updated_time" ` //
	Models             string  `json:"models"               orm:"models"               ` //
	Group              string  `json:"group"                orm:"group"                ` //
	UsedQuota          int64   `json:"used_quota"           orm:"used_quota"           ` //
	ModelMapping       string  `json:"model_mapping"        orm:"model_mapping"        ` //
	Priority           int64   `json:"priority"             orm:"priority"             ` //
	Config             string  `json:"config"               orm:"config"               ` //
	SystemPrompt       string  `json:"system_prompt"        orm:"system_prompt"        ` //
}
