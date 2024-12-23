//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/internal"
	"github.com/energye/lcl/api/internal/lcl"
)

var (
	releaseCallback    func()
	canWidgetSetInit   bool
	customWidgetImport *imports.Imports // LCL 自定义组件 interfaces 初始化导入
	liblclPreDefImport *imports.Imports // LCL 预定义导入
)

func init() {
	customWidgetImport = new(imports.Imports) // 自定义组件初始化导入
	liblclPreDefImport = new(imports.Imports) // LCL 部分预定义导入
}

// NewDefaultImports 创建默认导入表调用
func NewDefaultImports() *imports.Imports {
	if uiLib == 0 {
		uiLib = loadUILib()
	}
	result := new(imports.Imports)
	result.SetDll(uiLib)
	result.SetOk(true)
	return result
}

// APIInit API初始化
func APIInit() {
	// 加载动态链接库
	uiLib = loadUILib()

	customWidgetImport.SetDll(uiLib)
	internal.InitCustomWidgetImport(customWidgetImport)

	liblclPreDefImport.SetDll(uiLib)
	lcl.InitPreDefsImport(liblclPreDefImport)

}

// LibRelease app run end
func LibRelease() {
	if releaseCallback != nil {
		releaseCallback()
	}
	CustomWidgetSetFinalization()
	// 开启了finalizerOn选项后，以防止关闭库后GC还没开始调用。
	callGC()
	// 运行结束后就结束close掉lib，不然他不会关掉的
	closeLib()
}

// SetReleaseCallback 应用运行结束后释放资源之前执行
func SetReleaseCallback(fn func()) {
	if releaseCallback == nil {
		releaseCallback = fn
	}
}

// LCLPreDef 预定义LcL导入表
func LCLPreDef() imports.CallImport {
	return liblclPreDefImport
}
