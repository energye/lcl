// ----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License 2.0
//
// ----------------------------------------

//go:build linux
// +build linux

package api

import (
	"github.com/energye/lcl/api/imports"
	"sync"
)

// WidgetSetInitialization
// 自定义组件初始化 Linux
// 在 Linux 仅 CEF 使用 GTK2 GTK3 初始化时使用
func WidgetSetInitialization() {
	interfaceWidgetAPI().SysCallN(0)
}

// WidgetSetFinalization
// 在 Linux 仅 CEF 使用 GTK2 GTK3 时使用
func WidgetSetFinalization() {
	interfaceWidgetAPI().SysCallN(1)
}

var (
	interfaceWidgetOnce   sync.Once
	interfaceWidgetImport *imports.Imports
)

func interfaceWidgetAPI() *imports.Imports {
	interfaceWidgetOnce.Do(func() {
		interfaceWidgetImport = NewDefaultImports()
		interfaceWidgetImport.Table = []*imports.Table{
			imports.NewTable("Custom_WidgetSetInitialization", 0),
			imports.NewTable("Custom_WidgetSetFinalization", 0),
		}
	})
	return interfaceWidgetImport
}
