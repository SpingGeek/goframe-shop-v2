// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCouponInfoDao is the data access object for the table user_coupon_info.
type UserCouponInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns UserCouponInfoColumns // columns contains all the column names of Table for convenient usage.
}

// UserCouponInfoColumns defines and stores column names for the table user_coupon_info.
type UserCouponInfoColumns struct {
	Id        string // 用户优惠券表
	UserId    string //
	CouponId  string //
	Status    string // 状态：1可用 2已用 3过期
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// userCouponInfoColumns holds the columns for the table user_coupon_info.
var userCouponInfoColumns = UserCouponInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	CouponId:  "coupon_id",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewUserCouponInfoDao creates and returns a new DAO object for table data access.
func NewUserCouponInfoDao() *UserCouponInfoDao {
	return &UserCouponInfoDao{
		group:   "default",
		table:   "user_coupon_info",
		columns: userCouponInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserCouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserCouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserCouponInfoDao) Columns() UserCouponInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserCouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserCouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserCouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
