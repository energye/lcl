//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// Release 用于接口形式返回的指针释放
// 该函数不适用于 TClass 类型
func Release(data uintptr) {
	if data == 0 {
		return
	}
	interfacedObjectPreDefAPI().SysCallN(0, data)
}

var (
	interfacedObjectPreDefOnce   base.Once
	interfacedObjectPreDefImport *imports.Imports = nil
)

func interfacedObjectPreDefAPI() *imports.Imports {
	interfacedObjectPreDefOnce.Do(func() {
		interfacedObjectPreDefImport = api.NewDefaultImports()
		interfacedObjectPreDefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("DRelease", 0),
		}
	})
	return interfacedObjectPreDefImport
}
