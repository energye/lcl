//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import "github.com/energye/lcl/api"

// 初始化
func RtlInit() {
	// 初始
	api.DSysLocale(&SysLocale)
}
