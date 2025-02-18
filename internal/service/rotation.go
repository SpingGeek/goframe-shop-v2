// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IRotation interface {
		Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error)
	}
)

var (
	localRotation IRotation
)

func Rotation() IRotation {
	if localRotation == nil {
		panic("implement not found for interface IRotation, forgot register?")
	}
	return localRotation
}

func RegisterRotation(i IRotation) {
	localRotation = i
}
