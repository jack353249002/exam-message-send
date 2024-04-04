// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CoMessageDao is the data access object for table co_message.
type CoMessageDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns CoMessageColumns // columns contains all the column names of Table for convenient usage.
}

// CoMessageColumns defines and stores column names for table co_message.
type CoMessageColumns struct {
	Id        string //
	Title     string //
	Body      string //
	Attach    string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// coMessageColumns holds the columns for table co_message.
var coMessageColumns = CoMessageColumns{
	Id:        "id",
	Title:     "title",
	Body:      "body",
	Attach:    "attach",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCoMessageDao creates and returns a new DAO object for table data access.
func NewCoMessageDao(proxy ...dao_interface.IDao) *CoMessageDao {
	var dao *CoMessageDao
	if len(proxy) > 0 {
		dao = &CoMessageDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: coMessageColumns,
		}
		return dao
	}

	return &CoMessageDao{
		group:   "default",
		table:   "co_message",
		columns: coMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CoMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CoMessageDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *CoMessageDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *CoMessageDao) Columns() CoMessageColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CoMessageDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *CoMessageDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CoMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
