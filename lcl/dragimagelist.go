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

// IDragImageList Parent: ICustomImageList
type IDragImageList interface {
	ICustomImageList
	BeginDrag(window types.HWND, X int32, Y int32) bool                                   // function
	DragLock(window types.HWND, xPos int32, yPos int32) bool                              // function
	DragMove(X int32, Y int32) bool                                                       // function
	EndDrag() bool                                                                        // function
	SetDragImage(index int32, hotSpotX int32, hotSpotY int32) bool                        // function
	DragUnlock()                                                                          // procedure
	HideDragImage()                                                                       // procedure
	ShowDragImage()                                                                       // procedure
	DragCursor() types.TCursor                                                            // property DragCursor Getter
	SetDragCursor(value types.TCursor)                                                    // property DragCursor Setter
	DragHotspot() types.TPoint                                                            // property DragHotspot Getter
	SetDragHotspot(value types.TPoint)                                                    // property DragHotspot Setter
	Dragging() bool                                                                       // property Dragging Getter
	DraggingResolution() IDragImageListResolution                                         // property DraggingResolution Getter
	ResolutionWithIntToDragImageListResolution(imageWidth int32) IDragImageListResolution // property Resolution Getter
}

type TDragImageList struct {
	TCustomImageList
}

func (m *TDragImageList) BeginDrag(window types.HWND, X int32, Y int32) bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListAPI().SysCallN(2, m.Instance(), uintptr(window), uintptr(X), uintptr(Y))
	return api.GoBool(r)
}

func (m *TDragImageList) DragLock(window types.HWND, xPos int32, yPos int32) bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListAPI().SysCallN(3, m.Instance(), uintptr(window), uintptr(xPos), uintptr(yPos))
	return api.GoBool(r)
}

func (m *TDragImageList) DragMove(X int32, Y int32) bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListAPI().SysCallN(4, m.Instance(), uintptr(X), uintptr(Y))
	return api.GoBool(r)
}

func (m *TDragImageList) EndDrag() bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TDragImageList) SetDragImage(index int32, hotSpotX int32, hotSpotY int32) bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListAPI().SysCallN(6, m.Instance(), uintptr(index), uintptr(hotSpotX), uintptr(hotSpotY))
	return api.GoBool(r)
}

func (m *TDragImageList) DragUnlock() {
	if !m.IsValid() {
		return
	}
	dragImageListAPI().SysCallN(7, m.Instance())
}

func (m *TDragImageList) HideDragImage() {
	if !m.IsValid() {
		return
	}
	dragImageListAPI().SysCallN(8, m.Instance())
}

func (m *TDragImageList) ShowDragImage() {
	if !m.IsValid() {
		return
	}
	dragImageListAPI().SysCallN(9, m.Instance())
}

func (m *TDragImageList) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := dragImageListAPI().SysCallN(10, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDragImageList) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	dragImageListAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TDragImageList) DragHotspot() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragImageListAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDragImageList) SetDragHotspot(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragImageListAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDragImageList) Dragging() bool {
	if !m.IsValid() {
		return false
	}
	r := dragImageListAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TDragImageList) DraggingResolution() IDragImageListResolution {
	if !m.IsValid() {
		return nil
	}
	r := dragImageListAPI().SysCallN(13, m.Instance())
	return AsDragImageListResolution(r)
}

func (m *TDragImageList) ResolutionWithIntToDragImageListResolution(imageWidth int32) IDragImageListResolution {
	if !m.IsValid() {
		return nil
	}
	r := dragImageListAPI().SysCallN(14, m.Instance(), uintptr(imageWidth))
	return AsDragImageListResolution(r)
}

// NewDragImageList class constructor
func NewDragImageList(owner IComponent) IDragImageList {
	r := dragImageListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsDragImageList(r)
}

// NewDragImageListSize class constructor
func NewDragImageListSize(width int32, height int32) IDragImageList {
	r := dragImageListAPI().SysCallN(1, uintptr(width), uintptr(height))
	return AsDragImageList(r)
}

func TDragImageListClass() types.TClass {
	r := dragImageListAPI().SysCallN(15)
	return types.TClass(r)
}

var (
	dragImageListOnce   base.Once
	dragImageListImport *imports.Imports = nil
)

func dragImageListAPI() *imports.Imports {
	dragImageListOnce.Do(func() {
		dragImageListImport = api.NewDefaultImports()
		dragImageListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDragImageList_Create", 0), // constructor NewDragImageList
			/* 1 */ imports.NewTable("TDragImageList_CreateSize", 0), // constructor NewDragImageListSize
			/* 2 */ imports.NewTable("TDragImageList_BeginDrag", 0), // function BeginDrag
			/* 3 */ imports.NewTable("TDragImageList_DragLock", 0), // function DragLock
			/* 4 */ imports.NewTable("TDragImageList_DragMove", 0), // function DragMove
			/* 5 */ imports.NewTable("TDragImageList_EndDrag", 0), // function EndDrag
			/* 6 */ imports.NewTable("TDragImageList_SetDragImage", 0), // function SetDragImage
			/* 7 */ imports.NewTable("TDragImageList_DragUnlock", 0), // procedure DragUnlock
			/* 8 */ imports.NewTable("TDragImageList_HideDragImage", 0), // procedure HideDragImage
			/* 9 */ imports.NewTable("TDragImageList_ShowDragImage", 0), // procedure ShowDragImage
			/* 10 */ imports.NewTable("TDragImageList_DragCursor", 0), // property DragCursor
			/* 11 */ imports.NewTable("TDragImageList_DragHotspot", 0), // property DragHotspot
			/* 12 */ imports.NewTable("TDragImageList_Dragging", 0), // property Dragging
			/* 13 */ imports.NewTable("TDragImageList_DraggingResolution", 0), // property DraggingResolution
			/* 14 */ imports.NewTable("TDragImageList_ResolutionWithIntToDragImageListResolution", 0), // property ResolutionWithIntToDragImageListResolution
			/* 15 */ imports.NewTable("TDragImageList_TClass", 0), // function TDragImageListClass
		}
	})
	return dragImageListImport
}
