// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCouponInfo is the golang structure for table user_coupon_info.
type UserCouponInfo struct {
	Id        int         `json:"id"        orm:"id"         description:"用户优惠券表"`
	UserId    int         `json:"userId"    orm:"user_id"    description:""`
	CouponId  int         `json:"couponId"  orm:"coupon_id"  description:""`
	Status    int         `json:"status"    orm:"status"     description:"状态：1可用 2已用 3过期"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
