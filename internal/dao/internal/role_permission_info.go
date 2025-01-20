// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RolePermissionInfoDao is the data access object for the table role_permission_info.
type RolePermissionInfoDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns RolePermissionInfoColumns // columns contains all the column names of Table for convenient usage.
}

// RolePermissionInfoColumns defines and stores column names for the table role_permission_info.
type RolePermissionInfoColumns struct {
	Id           string //
	RoleId       string // 角色id
	PermissionId string // 权限id
	CreatedAt    string //
	UpdatedAt    string //
}

// rolePermissionInfoColumns holds the columns for the table role_permission_info.
var rolePermissionInfoColumns = RolePermissionInfoColumns{
	Id:           "id",
	RoleId:       "role_id",
	PermissionId: "permission_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewRolePermissionInfoDao creates and returns a new DAO object for table data access.
func NewRolePermissionInfoDao() *RolePermissionInfoDao {
	return &RolePermissionInfoDao{
		group:   "default",
		table:   "role_permission_info",
		columns: rolePermissionInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RolePermissionInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RolePermissionInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RolePermissionInfoDao) Columns() RolePermissionInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RolePermissionInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RolePermissionInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RolePermissionInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
