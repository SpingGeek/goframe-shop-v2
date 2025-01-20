// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"goframe-shop-v2/internal/dao/internal"
)

// internalPraiseInfoDao is an internal type for wrapping the internal DAO implementation.
type internalPraiseInfoDao = *internal.PraiseInfoDao

// praiseInfoDao is the data access object for the table praise_info.
// You can define custom methods on it to extend its functionality as needed.
type praiseInfoDao struct {
	internalPraiseInfoDao
}

var (
	// PraiseInfo is a globally accessible object for table praise_info operations.
	PraiseInfo = praiseInfoDao{
		internal.NewPraiseInfoDao(),
	}
)

// Add your custom methods and functionality below.
