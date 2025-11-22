//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package api

import (
	"github.com/energye/lcl/api/imports"
	"sync"
)

func NSWindow_FromForm(obj uintptr) uintptr {
	return platformPreDefAPI().SysCallN(0, obj)
}

func NSWindow_titleVisibility(n uintptr) uintptr {
	return platformPreDefAPI().SysCallN(1, n)
}

func NSWindow_setTitleVisibility(n uintptr, flag uintptr) {
	platformPreDefAPI().SysCallN(2, n, flag)
}

func NSWindow_titlebarAppearsTransparent(n uintptr) bool {
	return GoBool(platformPreDefAPI().SysCallN(3, n))
}

func NSWindow_setTitlebarAppearsTransparent(n uintptr, flag bool) {
	platformPreDefAPI().SysCallN(4, n, PasBool(flag))
}

func NSWindow_styleMask(n uintptr) uint {
	return uint(platformPreDefAPI().SysCallN(5, n))
}

func NSWindow_setStyleMask(n uintptr, mask uint) {
	platformPreDefAPI().SysCallN(6, n, uintptr(mask))
}

func NSWindow_setRepresentedURL(n uintptr, url uintptr) {
	platformPreDefAPI().SysCallN(7, n, url)
}

var (
	libPlatformPreDefOnce   sync.Once
	libPlatformPreDefImport *imports.Imports
)

func platformPreDefAPI() *imports.Imports {
	libPlatformPreDefOnce.Do(func() {
		libPlatformPreDefImport = NewDefaultImports()
		libPlatformPreDefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("NSWindow_FromForm", 0),
			/* 1 */ imports.NewTable("NSWindow_titleVisibility", 0),
			/* 2 */ imports.NewTable("NSWindow_setTitleVisibility", 0),
			/* 3 */ imports.NewTable("NSWindow_titlebarAppearsTransparent", 0),
			/* 4 */ imports.NewTable("NSWindow_setTitlebarAppearsTransparent", 0),
			/* 5 */ imports.NewTable("NSWindow_styleMask", 0),
			/* 6 */ imports.NewTable("NSWindow_setStyleMask", 0),
			/* 7 */ imports.NewTable("NSWindow_setRepresentedURL", 0),
		}
	})
	return libPlatformPreDefImport
}
