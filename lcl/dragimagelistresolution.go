//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// IDragImageListResolution Parent: ICustomImageListResolution
type IDragImageListResolution interface {
	ICustomImageListResolution
	BeginDrag(window types.HWND, X int32, Y int32) bool      // function
	DragLock(window types.HWND, xPos int32, yPos int32) bool // function
	DragMove(X int32, Y int32) bool                          // function
	EndDrag() bool                                           // function
	DragUnlock()                                             // procedure
	HideDragImage()                                          // procedure
	ShowDragImage()                                          // procedure
	DragHotspot() types.TPoint                               // property DragHotspot Getter
	SetDragHotspot(value types.TPoint)                       // property DragHotspot Setter
	Dragging() bool                                          // property Dragging Getter
}

type TDragImageListResolution struct {
	TCustomImageListResolution
}

func (m *TDragImageListResolution) BeginDrag(window types.HWND, X int32, Y int32) bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListResolutionAPI().SysCallN(1, m.Instance(), uintptr(window), uintptr(X), uintptr(Y))
	return api.GoBool(r)
}

func (m *TDragImageListResolution) DragLock(window types.HWND, xPos int32, yPos int32) bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListResolutionAPI().SysCallN(2, m.Instance(), uintptr(window), uintptr(xPos), uintptr(yPos))
	return api.GoBool(r)
}

func (m *TDragImageListResolution) DragMove(X int32, Y int32) bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListResolutionAPI().SysCallN(3, m.Instance(), uintptr(X), uintptr(Y))
	return api.GoBool(r)
}

func (m *TDragImageListResolution) EndDrag() bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListResolutionAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TDragImageListResolution) DragUnlock() {
	if !m.IsValid() {
		return
	}
	dragImageListResolutionAPI().SysCallN(5, m.Instance())
}

func (m *TDragImageListResolution) HideDragImage() {
	if !m.IsValid() {
		return
	}
	dragImageListResolutionAPI().SysCallN(6, m.Instance())
}

func (m *TDragImageListResolution) ShowDragImage() {
	if !m.IsValid() {
		return
	}
	dragImageListResolutionAPI().SysCallN(7, m.Instance())
}

func (m *TDragImageListResolution) DragHotspot() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragImageListResolutionAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDragImageListResolution) SetDragHotspot(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragImageListResolutionAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDragImageListResolution) Dragging() bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListResolutionAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

// NewDragImageListResolution class constructor
func NewDragImageListResolution(theOwner IComponent) IDragImageListResolution {
	r := dragImageListResolutionAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsDragImageListResolution(r)
}

func TDragImageListResolutionClass() types.TClass {
	r := dragImageListResolutionAPI().SysCallN(10)
	return types.TClass(r)
}

var (
	dragImageListResolutionOnce   base.Once
	dragImageListResolutionImport *imports.Imports = nil
)

func dragImageListResolutionAPI() *imports.Imports {
	dragImageListResolutionOnce.Do(func() {
		dragImageListResolutionImport = api.NewDefaultImports()
		dragImageListResolutionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDragImageListResolution_Create", 0), // constructor NewDragImageListResolution
			/* 1 */ imports.NewTable("TDragImageListResolution_BeginDrag", 0), // function BeginDrag
			/* 2 */ imports.NewTable("TDragImageListResolution_DragLock", 0), // function DragLock
			/* 3 */ imports.NewTable("TDragImageListResolution_DragMove", 0), // function DragMove
			/* 4 */ imports.NewTable("TDragImageListResolution_EndDrag", 0), // function EndDrag
			/* 5 */ imports.NewTable("TDragImageListResolution_DragUnlock", 0), // procedure DragUnlock
			/* 6 */ imports.NewTable("TDragImageListResolution_HideDragImage", 0), // procedure HideDragImage
			/* 7 */ imports.NewTable("TDragImageListResolution_ShowDragImage", 0), // procedure ShowDragImage
			/* 8 */ imports.NewTable("TDragImageListResolution_DragHotspot", 0), // property DragHotspot
			/* 9 */ imports.NewTable("TDragImageListResolution_Dragging", 0), // property Dragging
			/* 10 */ imports.NewTable("TDragImageListResolution_TClass", 0), // function TDragImageListResolutionClass
		}
	})
	return dragImageListResolutionImport
}
