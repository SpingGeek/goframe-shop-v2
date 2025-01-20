// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectionInfoDao is the data access object for the table collection_info.
type CollectionInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns CollectionInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CollectionInfoColumns defines and stores column names for the table collection_info.
type CollectionInfoColumns struct {
	Id        string //
	UserId    string // 用户id
	ObjectId  string // 对象id
	Type      string // 收藏类型：1商品 2文章
	CreatedAt string //
	UpdatedAt string //
}

// collectionInfoColumns holds the columns for the table collection_info.
var collectionInfoColumns = CollectionInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	ObjectId:  "object_id",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCollectionInfoDao creates and returns a new DAO object for table data access.
func NewCollectionInfoDao() *CollectionInfoDao {
	return &CollectionInfoDao{
		group:   "default",
		table:   "collection_info",
		columns: collectionInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CollectionInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CollectionInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CollectionInfoDao) Columns() CollectionInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CollectionInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CollectionInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CollectionInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
