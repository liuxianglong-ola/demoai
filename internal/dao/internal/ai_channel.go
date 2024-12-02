// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AiChannelDao is the data access object for table ai_channel.
type AiChannelDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AiChannelColumns // columns contains all the column names of Table for convenient usage.
}

// AiChannelColumns defines and stores column names for table ai_channel.
type AiChannelColumns struct {
	Id           string // 自增id
	Key          string // 密钥
	Status       string // 状态
	Name         string // 名称
	CreatedTime  string // 创建时间
	TestTime     string // 测试时间
	ResponseTime string // 测试响应时间
	BaseUrl      string // 代理地址
	Other        string //
	ModelStr     string // 具体的模型
	Group        string // 组
	Config       string // 配置项
}

// aiChannelColumns holds the columns for table ai_channel.
var aiChannelColumns = AiChannelColumns{
	Id:           "id",
	Key:          "key",
	Status:       "status",
	Name:         "name",
	CreatedTime:  "created_time",
	TestTime:     "test_time",
	ResponseTime: "response_time",
	BaseUrl:      "base_url",
	Other:        "other",
	ModelStr:     "model_str",
	Group:        "group",
	Config:       "config",
}

// NewAiChannelDao creates and returns a new DAO object for table data access.
func NewAiChannelDao() *AiChannelDao {
	return &AiChannelDao{
		group:   "default",
		table:   "ai_channel",
		columns: aiChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AiChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AiChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AiChannelDao) Columns() AiChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AiChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AiChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AiChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
