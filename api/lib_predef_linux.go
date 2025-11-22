//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux
// +build linux

package api

import (
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/types"
	"sync"
	"unsafe"
)

type TGdkWindowEdge int32

const (
	// drag window resize window
	GDK_WINDOW_EDGE_NORTH_WEST TGdkWindowEdge = iota
	GDK_WINDOW_EDGE_NORTH
	GDK_WINDOW_EDGE_NORTH_EAST
	GDK_WINDOW_EDGE_WEST
	GDK_WINDOW_EDGE_EAST
	GDK_WINDOW_EDGE_SOUTH_WEST
	GDK_WINDOW_EDGE_SOUTH
	GDK_WINDOW_EDGE_SOUTH_EAST
)

func GdkWindow_GetXId(g uintptr) (res uintptr) {
	platformPreDefAPI().SysCallN(0, g, uintptr(unsafe.Pointer(&res)))
	return
}

func GtkWidget_GetGtkFixed(g uintptr) uintptr {
	return platformPreDefAPI().SysCallN(1, g)
}

func GdkWindow_FromForm(obj uintptr) uintptr {
	return platformPreDefAPI().SysCallN(2, obj)
}

func GtkWidget_Window(g uintptr) uintptr {
	return platformPreDefAPI().SysCallN(3, g)
}

func GtkWindow_Drag(handle types.HWND, windowLeft, windowTop, button int32, edge TGdkWindowEdge) {
	platformPreDefAPI().SysCallN(4, uintptr(handle), uintptr(windowLeft), uintptr(windowTop), uintptr(button), uintptr(edge))
}

func Gtk3Widget_GtkWindow(handle types.HWND) uintptr {
	return platformPreDefAPI().SysCallN(5, handle)
}

func Gtk3Widget_Widget(handle types.HWND) uintptr {
	return platformPreDefAPI().SysCallN(6, handle)
}

var (
	libPlatformPreDefOnce   sync.Once
	libPlatformPreDefImport *imports.Imports
)

func platformPreDefAPI() *imports.Imports {
	libPlatformPreDefOnce.Do(func() {
		libPlatformPreDefImport = NewDefaultImports()
		libPlatformPreDefImport.Table = []*imports.Table{
			imports.NewTable("GdkWindow_GetXId", 0),
			imports.NewTable("GtkWidget_GetGtkFixed", 0),
			imports.NewTable("GdkWindow_FromForm", 0),
			imports.NewTable("GtkWidget_Window", 0),
			imports.NewTable("GtkWindow_Drag", 0),
			imports.NewTable("Gtk3Widget_GtkWindow", 0),
			imports.NewTable("Gtk3Widget_Widget", 0),
		}
	})
	return libPlatformPreDefImport
}
