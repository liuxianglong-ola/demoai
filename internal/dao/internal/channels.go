// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChannelsDao is the data access object for table channels.
type ChannelsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ChannelsColumns // columns contains all the column names of Table for convenient usage.
}

// ChannelsColumns defines and stores column names for table channels.
type ChannelsColumns struct {
	Id                 string //
	Type               string //
	Key                string //
	Status             string //
	Name               string //
	Weight             string //
	CreatedTime        string //
	TestTime           string //
	ResponseTime       string //
	BaseUrl            string //
	Other              string //
	Balance            string //
	BalanceUpdatedTime string //
	Models             string //
	Group              string //
	UsedQuota          string //
	ModelMapping       string //
	Priority           string //
	Config             string //
	SystemPrompt       string //
}

// channelsColumns holds the columns for table channels.
var channelsColumns = ChannelsColumns{
	Id:                 "id",
	Type:               "type",
	Key:                "key",
	Status:             "status",
	Name:               "name",
	Weight:             "weight",
	CreatedTime:        "created_time",
	TestTime:           "test_time",
	ResponseTime:       "response_time",
	BaseUrl:            "base_url",
	Other:              "other",
	Balance:            "balance",
	BalanceUpdatedTime: "balance_updated_time",
	Models:             "models",
	Group:              "group",
	UsedQuota:          "used_quota",
	ModelMapping:       "model_mapping",
	Priority:           "priority",
	Config:             "config",
	SystemPrompt:       "system_prompt",
}

// NewChannelsDao creates and returns a new DAO object for table data access.
func NewChannelsDao() *ChannelsDao {
	return &ChannelsDao{
		group:   "default",
		table:   "channels",
		columns: channelsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChannelsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChannelsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChannelsDao) Columns() ChannelsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChannelsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChannelsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChannelsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
