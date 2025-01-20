// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	Name         string      `json:"name"         orm:"name"          description:"用户名"`
	Avatar       string      `json:"avatar"       orm:"avatar"        description:"头像"`
	Password     string      `json:"password"     orm:"password"      description:""`
	UserSalt     string      `json:"userSalt"     orm:"user_salt"     description:"加密盐 生成密码用"`
	Sex          int         `json:"sex"          orm:"sex"           description:"1男 2女"`
	Status       int         `json:"status"       orm:"status"        description:"1正常 2拉黑冻结"`
	Sign         string      `json:"sign"         orm:"sign"          description:"个性签名"`
	SecretAnswer string      `json:"secretAnswer" orm:"secret_answer" description:"密保问题的答案"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`
}
