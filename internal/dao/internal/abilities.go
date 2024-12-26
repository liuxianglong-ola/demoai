// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AbilitiesDao is the data access object for table abilities.
type AbilitiesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AbilitiesColumns // columns contains all the column names of Table for convenient usage.
}

// AbilitiesColumns defines and stores column names for table abilities.
type AbilitiesColumns struct {
	Group     string //
	Model     string //
	ChannelId string //
	Enabled   string //
	Priority  string //
}

// abilitiesColumns holds the columns for table abilities.
var abilitiesColumns = AbilitiesColumns{
	Group:     "group",
	Model:     "model",
	ChannelId: "channel_id",
	Enabled:   "enabled",
	Priority:  "priority",
}

// NewAbilitiesDao creates and returns a new DAO object for table data access.
func NewAbilitiesDao() *AbilitiesDao {
	return &AbilitiesDao{
		group:   "default",
		table:   "abilities",
		columns: abilitiesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AbilitiesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AbilitiesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AbilitiesDao) Columns() AbilitiesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AbilitiesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AbilitiesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AbilitiesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
