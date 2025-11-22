// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

// Package Application environment initialization loads dynamic libraries

package initialize

import (
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/rtl/version"
)

// Initialize
// 初始化，运行时加载 LibLCL
func Initialize(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	// 加载 lib dll
	loadLibLCL(libs, resources)
	// 内嵌资源
	emfs.SetEMFS(libs, resources)
	// rtl
	rtl.Init()
	// version
	version.Init()
	// system platform api
	APIInit()
}
