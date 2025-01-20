// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RotationInfoDao is the data access object for the table rotation_info.
type RotationInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns RotationInfoColumns // columns contains all the column names of Table for convenient usage.
}

// RotationInfoColumns defines and stores column names for the table rotation_info.
type RotationInfoColumns struct {
	Id        string //
	PicUrl    string // 轮播图片
	Link      string // 跳转链接
	Sort      string // 排序字段
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// rotationInfoColumns holds the columns for the table rotation_info.
var rotationInfoColumns = RotationInfoColumns{
	Id:        "id",
	PicUrl:    "pic_url",
	Link:      "link",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewRotationInfoDao creates and returns a new DAO object for table data access.
func NewRotationInfoDao() *RotationInfoDao {
	return &RotationInfoDao{
		group:   "default",
		table:   "rotation_info",
		columns: rotationInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RotationInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RotationInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RotationInfoDao) Columns() RotationInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RotationInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RotationInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RotationInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
