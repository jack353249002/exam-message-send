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

// SysPermissionDao is the data access object for table sys_permission.
type SysPermissionDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysPermissionColumns // columns contains all the column names of Table for convenient usage.
}

// SysPermissionColumns defines and stores column names for table sys_permission.
type SysPermissionColumns struct {
	Id          string // ID
	ParentId    string // 父级ID
	Name        string // 名称
	Description string // 描述
	Identifier  string // 标识符
	Type        string // 类型：1api，2menu
	MatchMode   string // 匹配模式：ID：0，标识符：1
	IsShow      string // 是否显示：0不显示 1显示
	Sort        string // 排序
	CreatedAt   string //
	UpdatedAt   string //
}

// sysPermissionColumns holds the columns for table sys_permission.
var sysPermissionColumns = SysPermissionColumns{
	Id:          "id",
	ParentId:    "parent_id",
	Name:        "name",
	Description: "description",
	Identifier:  "identifier",
	Type:        "type",
	MatchMode:   "match_mode",
	IsShow:      "is_show",
	Sort:        "sort",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSysPermissionDao creates and returns a new DAO object for table data access.
func NewSysPermissionDao(proxy ...dao_interface.IDao) *SysPermissionDao {
	var dao *SysPermissionDao
	if len(proxy) > 0 {
		dao = &SysPermissionDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: sysPermissionColumns,
		}
		return dao
	}

	return &SysPermissionDao{
		group:   "default",
		table:   "sys_permission",
		columns: sysPermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPermissionDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPermissionDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysPermissionDao) Columns() SysPermissionColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPermissionDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SysPermissionDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
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
func (dao *SysPermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}