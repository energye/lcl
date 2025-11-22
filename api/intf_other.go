//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !linux
// +build !linux

package api

// WidgetSetInitialization
// 自定义组件初始化 Linux
// 在 Linux 仅 CEF 使用 GTK2 GTK3 初始化时使用
func WidgetSetInitialization() {
}

// WidgetSetFinalization
// 在 Linux 仅 CEF 使用 GTK2 GTK3 时使用
func WidgetSetFinalization() {
}
