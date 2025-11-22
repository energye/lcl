//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package misc

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

func DCClipRegionValid(dc types.HDC) bool {
	r1, _, _ := miscAPI().Proc(0).Call(dc)
	return api.GoBool(r1)
}

func DeallocateHWnd(wnd types.HWND) {
	miscAPI().Proc(1).Call(wnd)
}

func DestroyRubberBand(rubberBand types.HWND) {
	miscAPI().Proc(2).Call(rubberBand)
}

func DrawDefaultDockImage(oldRect, newRect types.TRect, operation types.TDockImageOperation) {
	miscAPI().Proc(3).Call(uintptr(base.UnsafePointer(&oldRect)), uintptr(base.UnsafePointer(&newRect)), uintptr(operation))
}

func DrawGrid(dc types.HDC, rect types.TRect, DX, DY int32) {
	miscAPI().Proc(4).Call(dc, uintptr(base.UnsafePointer(&rect)), uintptr(DX), uintptr(DY))
}

func Frame3d(dc types.HDC, rect *types.TRect, frameWidth int32, style types.TGraphicsBevelCut) bool {
	r1, _, _ := miscAPI().Proc(5).Call(dc, uintptr(base.UnsafePointer(rect)), uintptr(frameWidth), uintptr(style))
	return api.GoBool(r1)
}

var (
	miscOnce   base.Once
	miscImport *imports.Imports = nil
)

func miscAPI() *imports.Imports {
	miscOnce.Do(func() {
		miscImport = api.NewDefaultImports()
		miscImport.Table = []*imports.Table{
			imports.NewTable("LCLMisc_DCClipRegionValid", 0),
			imports.NewTable("LCLMisc_DeallocateHWnd", 0),
			imports.NewTable("LCLMisc_DestroyRubberBand", 0),
			imports.NewTable("LCLMisc_DrawDefaultDockImage", 0),
			imports.NewTable("LCLMisc_DrawGrid", 0),
			imports.NewTable("LCLMisc_Frame3d", 0),
		}
	})
	return miscImport
}
