// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleInfoDao is the data access object for the table role_info.
type RoleInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns RoleInfoColumns // columns contains all the column names of Table for convenient usage.
}

// RoleInfoColumns defines and stores column names for the table role_info.
type RoleInfoColumns struct {
	Id        string //
	Name      string // 角色名称
	Desc      string // 描述
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// roleInfoColumns holds the columns for the table role_info.
var roleInfoColumns = RoleInfoColumns{
	Id:        "id",
	Name:      "name",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewRoleInfoDao creates and returns a new DAO object for table data access.
func NewRoleInfoDao() *RoleInfoDao {
	return &RoleInfoDao{
		group:   "default",
		table:   "role_info",
		columns: roleInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RoleInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RoleInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RoleInfoDao) Columns() RoleInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RoleInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RoleInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RoleInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
