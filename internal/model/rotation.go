// Package model @Author Gopher
// @Date 2025/1/20 22:06:00
// @Desc
package model

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}
