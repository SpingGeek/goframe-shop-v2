// Package backend @Author Gopher
// @Date 2025/1/20 21:31:00
// @Desc
package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}

type RotationRes struct {
	//todo
	//g.Meta `mime:"text/html" example:"string"`
	RotationId int `json:"rotationId"`
}
