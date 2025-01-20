// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderGoodsInfoDao is the data access object for the table order_goods_info.
type OrderGoodsInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns OrderGoodsInfoColumns // columns contains all the column names of Table for convenient usage.
}

// OrderGoodsInfoColumns defines and stores column names for the table order_goods_info.
type OrderGoodsInfoColumns struct {
	Id             string // 商品维度的订单表
	OrderId        string // 关联的主订单表
	GoodsId        string // 商品id
	GoodsOptionsId string // 商品规格id sku id
	Count          string // 商品数量
	Remark         string // 备注
	Price          string // 订单金额 单位分
	CouponPrice    string // 优惠券金额 单位分
	ActualPrice    string // 实际支付金额 单位分
	CreatedAt      string //
	UpdatedAt      string //
}

// orderGoodsInfoColumns holds the columns for the table order_goods_info.
var orderGoodsInfoColumns = OrderGoodsInfoColumns{
	Id:             "id",
	OrderId:        "order_id",
	GoodsId:        "goods_id",
	GoodsOptionsId: "goods_options_id",
	Count:          "count",
	Remark:         "remark",
	Price:          "price",
	CouponPrice:    "coupon_price",
	ActualPrice:    "actual_price",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewOrderGoodsInfoDao creates and returns a new DAO object for table data access.
func NewOrderGoodsInfoDao() *OrderGoodsInfoDao {
	return &OrderGoodsInfoDao{
		group:   "default",
		table:   "order_goods_info",
		columns: orderGoodsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrderGoodsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderGoodsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrderGoodsInfoDao) Columns() OrderGoodsInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderGoodsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderGoodsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OrderGoodsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
