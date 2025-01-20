// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CartInfo is the golang structure for table cart_info.
type CartInfo struct {
	Id             int         `json:"id"             orm:"id"               description:"购物车表"`
	UserId         int         `json:"userId"         orm:"user_id"          description:""`
	GoodsOptionsId int         `json:"goodsOptionsId" orm:"goods_options_id" description:"商品规格id"`
	Count          int         `json:"count"          orm:"count"            description:"商品数量"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:""`
}
