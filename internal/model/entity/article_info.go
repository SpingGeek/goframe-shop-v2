// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleInfo is the golang structure for table article_info.
type ArticleInfo struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	UserId    int         `json:"userId"    orm:"user_id"    description:"作者id"`
	Title     string      `json:"title"     orm:"title"      description:"标题"`
	Desc      string      `json:"desc"      orm:"desc"       description:"摘要"`
	PicUrl    string      `json:"picUrl"    orm:"pic_url"    description:"封面图"`
	IsAdmin   int         `json:"isAdmin"   orm:"is_admin"   description:"1后台管理员发布 2前台用户发布"`
	Praise    int         `json:"praise"    orm:"praise"     description:"点赞数"`
	Detail    string      `json:"detail"    orm:"detail"     description:"文章详情"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
