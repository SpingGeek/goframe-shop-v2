// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminInfoDao is the data access object for the table admin_info.
type AdminInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns AdminInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AdminInfoColumns defines and stores column names for the table admin_info.
type AdminInfoColumns struct {
	Id        string //
	Name      string // 用户名
	Password  string // 密码
	RoleIds   string // 角色ids
	CreatedAt string //
	UpdatedAt string //
	UserSalt  string // 加密盐
	IsAdmin   string // 是否超级管理员
}

// adminInfoColumns holds the columns for the table admin_info.
var adminInfoColumns = AdminInfoColumns{
	Id:        "id",
	Name:      "name",
	Password:  "password",
	RoleIds:   "role_ids",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	UserSalt:  "user_salt",
	IsAdmin:   "is_admin",
}

// NewAdminInfoDao creates and returns a new DAO object for table data access.
func NewAdminInfoDao() *AdminInfoDao {
	return &AdminInfoDao{
		group:   "default",
		table:   "admin_info",
		columns: adminInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminInfoDao) Columns() AdminInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AdminInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
