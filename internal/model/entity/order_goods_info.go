// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsInfo is the golang structure for table order_goods_info.
type OrderGoodsInfo struct {
	Id             int         `json:"id"             orm:"id"               description:"商品维度的订单表"`
	OrderId        int         `json:"orderId"        orm:"order_id"         description:"关联的主订单表"`
	GoodsId        int         `json:"goodsId"        orm:"goods_id"         description:"商品id"`
	GoodsOptionsId int         `json:"goodsOptionsId" orm:"goods_options_id" description:"商品规格id sku id"`
	Count          int         `json:"count"          orm:"count"            description:"商品数量"`
	Remark         string      `json:"remark"         orm:"remark"           description:"备注"`
	Price          int         `json:"price"          orm:"price"            description:"订单金额 单位分"`
	CouponPrice    int         `json:"couponPrice"    orm:"coupon_price"     description:"优惠券金额 单位分"`
	ActualPrice    int         `json:"actualPrice"    orm:"actual_price"     description:"实际支付金额 单位分"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`
}
