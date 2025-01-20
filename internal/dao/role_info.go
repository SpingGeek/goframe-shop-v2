// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"goframe-shop-v2/internal/dao/internal"
)

// internalRoleInfoDao is an internal type for wrapping the internal DAO implementation.
type internalRoleInfoDao = *internal.RoleInfoDao

// roleInfoDao is the data access object for the table role_info.
// You can define custom methods on it to extend its functionality as needed.
type roleInfoDao struct {
	internalRoleInfoDao
}

var (
	// RoleInfo is a globally accessible object for table role_info operations.
	RoleInfo = roleInfoDao{
		internal.NewRoleInfoDao(),
	}
)

// Add your custom methods and functionality below.
