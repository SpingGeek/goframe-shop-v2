// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PositionInfoDao is the data access object for the table position_info.
type PositionInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns PositionInfoColumns // columns contains all the column names of Table for convenient usage.
}

// PositionInfoColumns defines and stores column names for the table position_info.
type PositionInfoColumns struct {
	Id        string //
	PicUrl    string // 图片链接
	GoodsName string // 商品名称
	Link      string // 跳转链接
	Sort      string // 排序
	GoodsId   string // 商品id
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// positionInfoColumns holds the columns for the table position_info.
var positionInfoColumns = PositionInfoColumns{
	Id:        "id",
	PicUrl:    "pic_url",
	GoodsName: "goods_name",
	Link:      "link",
	Sort:      "sort",
	GoodsId:   "goods_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPositionInfoDao creates and returns a new DAO object for table data access.
func NewPositionInfoDao() *PositionInfoDao {
	return &PositionInfoDao{
		group:   "default",
		table:   "position_info",
		columns: positionInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PositionInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PositionInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PositionInfoDao) Columns() PositionInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PositionInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PositionInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PositionInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
