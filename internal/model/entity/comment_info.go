// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentInfo is the golang structure for table comment_info.
type CommentInfo struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	ParentId  int         `json:"parentId"  orm:"parent_id"  description:"父级评论id"`
	UserId    int         `json:"userId"    orm:"user_id"    description:""`
	ObjectId  int         `json:"objectId"  orm:"object_id"  description:""`
	Type      int         `json:"type"      orm:"type"       description:"评论类型：1商品 2文章"`
	Content   string      `json:"content"   orm:"content"    description:"评论内容"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
