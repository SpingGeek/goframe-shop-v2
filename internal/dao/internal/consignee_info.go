// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsigneeInfoDao is the data access object for the table consignee_info.
type ConsigneeInfoDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns ConsigneeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// ConsigneeInfoColumns defines and stores column names for the table consignee_info.
type ConsigneeInfoColumns struct {
	Id        string // 收货地址表
	UserId    string //
	IsDefault string // 默认地址1  非默认0
	Name      string //
	Phone     string //
	Province  string //
	City      string //
	Town      string // 县区
	Street    string // 街道乡镇
	Detail    string // 地址详情
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// consigneeInfoColumns holds the columns for the table consignee_info.
var consigneeInfoColumns = ConsigneeInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	IsDefault: "is_default",
	Name:      "name",
	Phone:     "phone",
	Province:  "province",
	City:      "city",
	Town:      "town",
	Street:    "street",
	Detail:    "detail",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewConsigneeInfoDao creates and returns a new DAO object for table data access.
func NewConsigneeInfoDao() *ConsigneeInfoDao {
	return &ConsigneeInfoDao{
		group:   "default",
		table:   "consignee_info",
		columns: consigneeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ConsigneeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ConsigneeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ConsigneeInfoDao) Columns() ConsigneeInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ConsigneeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ConsigneeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ConsigneeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
