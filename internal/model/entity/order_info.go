// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	Id               int         `json:"id"               orm:"id"                description:""`
	Number           string      `json:"number"           orm:"number"            description:"订单编号"`
	UserId           int         `json:"userId"           orm:"user_id"           description:"用户id"`
	PayType          int         `json:"payType"          orm:"pay_type"          description:"支付方式 1微信 2支付宝 3云闪付"`
	Remark           string      `json:"remark"           orm:"remark"            description:"备注"`
	PayAt            *gtime.Time `json:"payAt"            orm:"pay_at"            description:"支付时间"`
	Status           int         `json:"status"           orm:"status"            description:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价"`
	ConsigneeName    string      `json:"consigneeName"    orm:"consignee_name"    description:"收货人姓名"`
	ConsigneePhone   string      `json:"consigneePhone"   orm:"consignee_phone"   description:"收货人手机号"`
	ConsigneeAddress string      `json:"consigneeAddress" orm:"consignee_address" description:"收货人详细地址"`
	Price            int         `json:"price"            orm:"price"             description:"订单金额 单位分"`
	CouponPrice      int         `json:"couponPrice"      orm:"coupon_price"      description:"优惠券金额 单位分"`
	ActualPrice      int         `json:"actualPrice"      orm:"actual_price"      description:"实际支付金额 单位分"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:""`
}
